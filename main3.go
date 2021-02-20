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

//Lista
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
	} else {
		ultimo := list.fin
		ultimo.siguiente = newNodo
		ultimo.siguiente.anterior = ultimo
		list.fin = newNodo
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

//funcion para matriz datos
func matriz(dat *Datos) {

	departamentos := make([]string, len(dat.Seccion[0].Departamento))

	for i := 0; i < len(dat.Seccion); i++ {
		indice[i] = dat.Seccion[i].Indice
	}

	for i := 0; i < (len(departamentos) - 1); i++ {
		departamentos[i] = dat.Seccion[i].Departamento[i].Nombre
	}

	fmt.Println("Esta aqui")
	fmt.Println(departamentos, indice)

	//Estableciendo tamaño al vector de la matriz linealizada
	Linealizacion = make([]ListaD, (len(dat.Seccion) * len(dat.Seccion) * len(dat.Seccion[0].Departamento) * 5))

	//uso de for por cada dimenasion de la matriz de datos
	for i := 0; i < len(dat.Seccion); i++ {
		for j := 0; j < len(dat.Seccion[i].Departamento); j++ {
			for k := 0; k < len(dat.Seccion[i].Departamento[j].Tiendas); k++ {
				//Uso de formual de row-major para linealizacion de matriz
				Linealizacion[(i*len(dat.Seccion[i].Departamento)+j)*5+(dat.Seccion[i].Departamento[j].Tiendas[k].Calificacion-1)].agregarTienda(Tienda{Nombre: dat.Seccion[i].Departamento[j].Tiendas[k].Nombre, Descripcion: dat.Seccion[i].Departamento[j].Tiendas[k].Descripcion, Contacto: dat.Seccion[i].Departamento[j].Tiendas[k].Contacto, Calificacion: dat.Seccion[i].Departamento[j].Tiendas[k].Calificacion})
			}
		}
	}

	//Imprimir matriz linalizada

	for i := 0; i < len(Linealizacion); i++ {
		fmt.Println(i)
		nodo := Linealizacion[i].inicio
		if nodo != nil {
			fmt.Println(nodo.tienda)
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

	tiendaE := buscarTienda(&paramBus)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	//Salida
	json.NewEncoder(w).Encode(tiendaE)

}

//Funcion para buscar la tienda en el vector linealizado
func buscarTienda(par *SearchShop) *Tienda {
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
				break
			}
		} else {
			return &shopF
		}

	}
	return &shopF
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
	fmt.Println(shop)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	//Salida
	json.NewEncoder(w).Encode(shop)
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
