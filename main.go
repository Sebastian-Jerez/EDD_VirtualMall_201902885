package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Lista Doble para alacenar tiendas------------------------------------
//Nodos
type Nodo struct {
	tienda    Tienda
	dato      string
	siguiente *Nodo
	anterior  *Nodo
}

//Lista doble
type ListaD struct {
	inicio *Nodo
	fin    *Nodo
	size   int
}

//Funciones de la lista doble----------
func (lista *ListaD) primero() *Nodo {
	return lista.inicio
}

func (lista *ListaD) ultimo() *Nodo {
	return lista.fin
}

func (node *Nodo) ant() *Nodo {
	return node.anterior
}

func (node *Nodo) sig() *Nodo {
	return node.siguiente
}

//Funcion para agregar nodo
func (list *ListaD) agregarTienda(store Tienda) {
	newNodo := &Nodo{tienda: store}

	if list.inicio == nil {
		list.inicio = newNodo
		list.fin = newNodo
		list.size++
	} else {
		ultimo := list.fin
		ultimo.siguiente = newNodo
		ultimo.siguiente.anterior = ultimo
		list.fin = newNodo
		list.size++s
	}
}

//Funcion para buscar nodo
func buscarN(name string, l *ListaD) *Nodo {
	encontrado := false
	var ret *Nodo = nil
	for n := l.primero(); n != nil && !encontrado; n = n.sig() {
		if n.tienda.Nombre == name {
			encontrado = true
			ret = n
		}
	}
	return ret
}

func paso(l *ListaD) {
	newN := l.inicio

	for newN.siguiente != nil {
		i := 0
		ListaTT = append(ListaTT, newN.tienda)
		newN = newN.siguiente
		fmt.Println(i)
		i++
	}
	fmt.Println(ListaTT)
}

//Structs para lectura de JSON--------------------------------
type Datos struct {
	Seccion []Seccion `json:"Datos"`
}

type Seccion struct {
	Indice       string          `json:"Indice"`
	Departamento []Departamentos `json:"Departamentos"`
}

type Departamentos struct {
	Nombre  string   `json:"Nombre"`
	Tiendas []Tienda `json:"Tiendas"`
}

type Tienda struct {
	Nombre       string `json:"Nombre"`
	Descripcion  string `json:"Descripcion"`
	Contacto     string `json:"Contacto"`
	Calificacion int    `json:"Calificacion"`
}

type SearchShop struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}

//vector de indices
var indice [27]string

//Vaector linealizado
var Linealizacion []ListaD

//vector de departamentod
var departamentos []string

//vecto de tiendas temporal
var ListaTT []Tienda

//funcion para matriz datos
func matriz(dat *Datos) {

	departamentos := make([]string, len(dat.Seccion[0].Departamento))

	for i := 0; i < len(dat.Seccion); i++ {
		indice[i] = dat.Seccion[i].Indice
	}

	for i := 0; i < (len(departamentos)); i++ {
		departamentos[i] = dat.Seccion[0].Departamento[i].Nombre
	}

	fmt.Println("Esta aqui")
	fmt.Println(departamentos, indice)

	//Estableciendo tamaño al vector de la matriz linealizada
	Linealizacion = make([]ListaD, (len(dat.Seccion) * len(dat.Seccion[0].Departamento) * 5))

	//uso de for por cada dimenasion de la matriz de datos
	for i := 0; i < len(dat.Seccion); i++ {
		for j := 0; j < len(dat.Seccion[i].Departamento); j++ {
			for k := 0; k < len(dat.Seccion[i].Departamento[j].Tiendas); k++ {
				//Uso de formual de row-major para linealizacion de matriz
				Linealizacion[(i*len(dat.Seccion[i].Departamento)+j)*5+(dat.Seccion[i].Departamento[j].Tiendas[k].Calificacion-1)].agregarTienda(dat.Seccion[i].Departamento[j].Tiendas[k])
			}
		}
	}

	//Imprimir matriz linealizada

	for i := 0; i < len(Linealizacion); i++ {
		fmt.Println(i)
		nodo := Linealizacion[i].inicio
		for nodo != nil {
			fmt.Println(nodo.tienda)
			nodo = nodo.siguiente
		}

	}
}

func getTiendas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(ListaTiendas)
}

//Ingreso del archivo json
func crearTienda(w http.ResponseWriter, r *http.Request) {
	var listaDat Datos
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Inserte datos validos")
	}

	json.Unmarshal(reqBody, &listaDat)

	matriz(&listaDat)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(listaDat)
}

func searchtienda(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//nombretienda := vars["nombre"]

	//for i, tienda := range ListaTiendas {
	//if tienda.Nombre == nombretienda {

	//}
	//}
}

func BusquedaPE(w http.ResponseWriter, r *http.Request) {
	var paramBus SearchShop
	//var tiendaE Tienda
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Inserte datos validos")
	}

	json.Unmarshal(reqBody, &paramBus)

	buscarTienda(&paramBus, w)

}

//Funcion para buscar la tienda en el vector linealizado
func buscarTienda(par *SearchShop, w http.ResponseWriter) {
	fmt.Println("Netra a buscarTienda")
	val := len(Linealizacion)
	fmt.Println(val)
	nombre := par.Nombre
	var shopF Tienda
	for i := 0; i < val; i++ {
		nodo := Linealizacion[i].inicio
		fmt.Println(i)
		if nodo != nil {
			fmt.Println("Entro a listaD")
			td := Linealizacion[i]
			if buscarN(nombre, &td) != nil {
				fmt.Println("LIsta")
				fmt.Println(td)
				fmt.Println(i)
				shopF = buscarN(nombre, &td).tienda
				//Salida
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(shopF)
				break
			}
		} else {

		}

	}
}

func deleteTienda(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//tiendaNombre := vars["nombre"]

	//for _, tienda := range ListaTiendas {
	//if tienda.Nombre == tiendaNombre {
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(tienda)
	//}
	//}
}

func BusquedaRL(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idTienda, err := strconv.Atoi(vars["num"])

	if err != nil {
		fmt.Errorf("Dato incorrecto")
	}
	fmt.Println(idTienda)
	shop := Linealizacion[idTienda]

	paso(&shop)
	fmt.Println(ListaTT)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	//Salida
	json.NewEncoder(w).Encode(ListaTT)
}

func grafo() {
	valD := 0

	txtdot := "digraph G { \nnode[shape=record]\n" + `graph[splines="ortho"]` + "\n"
	rank := "{rank=same;"

	//Obtener todos los indices
	for i := 0; i < len(indice); i++ {
		if indice[i] != "" {
			valD++
		}
	}
	indices := make([]string, valD)
	for i := 0; i < len(indice); i++ {
		if indice[i] != "" {
			indices = append(indices, indice[i])
		}
	}

	//Agrgar texto dot
	for i := 0; i < len(indices); i++ {
		for j := 0; j < len(departamentos); j++ {

			//Crear nodos por departamento
			for k := 0; i < 5; k++ {
				txtdot += "nodo" + strconv.Itoa(k) + `[label="` + indices[i] + "|" + departamentos[j] + "| Posicion:" + strconv.Itoa(k+1) + `"]` + "\n"
				rank += "nodo" + strconv.Itoa(k) + ";"

			}
			rank += "}\n"
			txtdot += rank
			rank = "{rank=same;"

			//indiceAc := indices[i]
			//departoAc := departamentos[j]

			//Crear conexiones entre nodos
			for k := 0; k < 4; k++ {
				txtdot += "nodo" + strconv.Itoa(k) + "->nodo" + strconv.Itoa(k+1) + "\n"
			}

			for i := 0; i < len(Linealizacion); i++ {
				//fmt.Println(i)
				nodo := Linealizacion[i].inicio
				if nodo != nil {
					cal := nodo.tienda.Calificacion
					tnd := Linealizacion[i]
					tnd.graphNodos(cal)

				}
			}

		}
	}

}

func (l *ListaD) graphNodos(n int) string {
	inicio := l.inicio
	nodos := "nodo" + strconv.Itoa(n) + "->"
	datos := ""
	if inicio != nil {
		for inicio != nil {
			datos += inicio.tienda.Nombre + `[label="` + inicio.tienda.Nombre + "|" + inicio.tienda.Contacto + "|" + strconv.Itoa(inicio.tienda.Calificacion) + `"]` + "\n"
			if inicio.siguiente != nil {
				nodos += inicio.tienda.Nombre + "->" + inicio.siguiente.tienda.Nombre
			}
			if l.inicio == l.fin {
				nodos += inicio.tienda.Nombre
			}
			inicio = inicio.siguiente
		}
		datos += nodos + "\n"
		return datos
	}
	return datos
}

func indexR(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor en funcionamiento")
}

func main() {
	fmt.Print("Hola")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexR)
	router.HandleFunc("/tiendas", getTiendas).Methods("GET")
	//router.HandleFunc("/tiendas", crearTienda).Methods("POST")
	router.HandleFunc("/cargartienda", crearTienda).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", BusquedaPE).Methods("POST")
	router.HandleFunc("/tiendas/{nombre}", searchtienda).Methods("GET")
	router.HandleFunc("/id/{num}", BusquedaRL).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
