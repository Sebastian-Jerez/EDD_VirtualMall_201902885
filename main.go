package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
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
		list.size++
	}
}

//Funcion para buscar y eliminar
func (l *ListaD) buscarYEN(name string) *Nodo {
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

//
func paso(l *ListaD) {
	newN := l.inicio

	for i := 0; i < l.size; i++ {
		ListaTT = append(ListaTT, newN.tienda)
		newN = newN.siguiente
	}
	fmt.Println(ListaTT)
}
func EliminarNodo(nombre string, l *ListaD) {

	nodoD := l.buscarYEN(nombre)
	if nodoD != nil {
		fmt.Println("NEtra para eliminar")
		if nodoD == l.inicio {
			l.inicio = nodoD.siguiente
			l.size--
		} else if nodoD == l.fin {
			l.fin = nodoD.anterior
			l.size--
		} else {
			nodoAnt := nodoD.anterior
			nodoSig := nodoD.siguiente
			// Remover el nodo
			nodoAnt.siguiente = nodoD.siguiente
			nodoSig.anterior = nodoD.anterior
			fmt.Println("Termina de eliminar")
			l.size--
		}

	}
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
	Logo         string `json:"Logo"`
}

type Pedidos struct {
	Compras []Compras `json:"Pedidos"`
}

type Compras struct {
	Fecha        string            `json:"Fecha"`
	Tienda       string            `json:"Tienda"`
	Departamento string            `json:"Departamento"`
	Calificacion int               `json:"Calificacion"`
	ProductosP   []ProductoPedidos `json:"Productos"`
}

type ProductoPedidos struct {
	Codigo int `json:"Codigo"`
}

type Inventarios struct {
	Inverntarios []Inverntario `json:"Invetarios"`
}

type Inverntario struct {
	Tienda       string                `json:"Tienda"`
	Departamento string                `json:"Departamento"`
	Calificacion int                   `json:"Calificacion"`
	ProductosI   []ProductosInventario `json:"Productos"`
}

type ProductosInventario struct {
	Nombre      string `json:"Nombre"`
	Codigo      int    `json:"Codigo"`
	Descripcion string `json:"Descripcion"`
	Precio      int    `json:"Precio"`
	Cantidad    int    `json:"Cantidad"`
	Imagen      string `json:"Imagen"`
}

//paramentros de busquea para buscar
type SearchShop struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}

//paramaetros para eliminar tienda
type DeleteShop struct {
	Nombre       string `json:"Nombre"`
	Categoria    string `json:"Categoria"`
	Calificacion int    `json:"Calificacion"`
}

//------------Estructuras para arbol AVL----------------------

//Nodos
type NodoAVL struct {
	Produ     ProductosInventario
	Izquierda *NodoAVL
	Derecha   *NodoAVL
	Balance   int
}

//Funcion para crear un nuevo nodo
func NewNodo(producto ProductosInventario) *NodoAVL {
	return &NodoAVL{producto, nil, nil, 0}
}

//Constantes para armar el AVL
const (
	Left_Heavy  = -1
	Balanced    = 0
	Right_Heavy = 1
)

//Arbol AVL
type Arbol struct {
	RaizAVL *NodoAVL
}

//Funcion para crer un nuevo arbol
func NewArbol() *Arbol {
	return &Arbol{nil}
}

//agregar un nuevo arbol
func ingresarA(arbol *Arbol, producto ProductosInventario) bool {
	increase := false

	return inseterarN(&arbol.RaizAVL, producto, &increase)
}

//Insertar un nuevoo nodo al arbol
func inseterarN(nodo **NodoAVL, producto ProductosInventario, increase *bool) bool {
	if *nodo == nil {
		*nodo = NewNodo(producto)
		*increase = true
		return true
	} else if (*nodo).Produ.Cantidad > producto.Cantidad {

		returnvalue := inseterarN(&(*nodo).Izquierda, producto, increase)

		if *increase {

			switch (*nodo).Balance {
			case Balanced:
				(*nodo).Balance = Left_Heavy
			case Right_Heavy:
				(*nodo).Balance = Balanced
				*increase = false
				break
			case Left_Heavy:
				RebalanceLeft(&*nodo)
				*increase = false
				break
			}
		}

		return returnvalue

	} else if (*nodo).Produ.Cantidad < producto.Cantidad {
		returnvalue2 := inseterarN(&(*nodo).Derecha, producto, increase)

		if *increase {
			switch (*nodo).Balance {
			case Balanced:
				(*nodo).Balance = Right_Heavy
			case Left_Heavy:
				(*nodo).Balance = Balanced
				*increase = false
				break
			case Right_Heavy:

			}
		}
		return returnvalue2
	}

	return false
}

//---------------Funciones de balanceo de Arbol AVL---------------------
//Funcion de balanceo por la izquierda Arbol AVL
func RebalanceLeft(raizL **NodoAVL) {
	hijoIz := (*raizL).Izquierda

	if hijoIz.Balance == Right_Heavy {
		hijoIzDer := hijoIz.Derecha

		if hijoIzDer.Balance == Left_Heavy {
			hijoIz.Balance = Balanced
			hijoIzDer.Balance = Balanced
			(*raizL).Balance = Right_Heavy
		} else if hijoIzDer.Balance == Balanced {
			hijoIz.Balance = Balanced
			hijoIzDer.Balance = Balanced
			(*raizL).Balance = Balanced
		} else {
			hijoIz.Balance = Left_Heavy
			hijoIzDer.Balance = Balanced
			(*raizL).Balance = Balanced
		}

		RotacionIzquierda(&(*raizL).Izquierda)
	} else {
		hijoIz.Balance = Balanced
		(*raizL).Balance = Balanced
	}

	RotacionDerecha(&*raizL)
}

//Funcion de balanceo por la derecha
func RebalanceRight(raizR **NodoAVL) {
	hijoDer := (*raizR).Derecha

	if hijoDer.Balance == Left_Heavy {
		hijoDerIz := hijoDer.Izquierda

		if hijoDerIz.Balance == Right_Heavy {
			hijoDer.Balance = Balanced
			hijoDerIz.Balance = Balanced
			(*raizR).Balance = Left_Heavy
		} else if hijoDerIz.Balance == Balanced {
			hijoDer.Balance = Balanced
			hijoDerIz.Balance = Balanced
			(*raizR).Balance = Balanced
		} else {
			hijoDer.Balance = Right_Heavy
			hijoDerIz.Balance = Balanced
			(*raizR).Balance = Balanced
		}
		RotacionDerecha(&(*raizR).Derecha)
	} else {
		hijoDer.Balance = Balanced
		(*raizR).Balance = Balanced
	}

	RotacionIzquierda(&*raizR)
}

//---------------------Funciones de rotacion------------------------
//Funcion de rotacion por la derecha
func RotacionDerecha(raizA **NodoAVL) {
	tmp := (*raizA).Izquierda
	(*raizA).Izquierda = tmp.Derecha
	tmp.Derecha = *raizA
	*raizA = tmp
}

//Funcion para rotacion por la izquierda
func RotacionIzquierda(raizA **NodoAVL) {
	tmp := (*raizA).Derecha
	(*raizA).Derecha = tmp.Izquierda
	tmp.Izquierda = *raizA
	*raizA = tmp
}

//-----------------Reportes del proyecto--------------------------------
//Grafo del arbol AVL
func GrafoAVL(arbol *Arbol) {
	inicioDot := "digraph G{\nnode [shape=circle];\n"

	dot := ""

	if arbol.RaizAVL != nil {
		RecorrerArbol(&arbol.RaizAVL, &dot)
	}

	inicioDot += dot + "\n}\n"

	path := "grafoAVL.dot"

	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if ExisteError(err) {
			return
		}

		defer file.Close()
		fmt.Println("Archivo creado con exito")
	}

	var file, err2 = os.OpenFile(path, os.O_RDWR, 0644)
	if ExisteError(err2) {
		return
	}

	defer file.Close()

	_, err = file.WriteString(inicioDot)
	if ExisteError(err) {
		return
	}

	err = file.Sync()
	if ExisteError(err) {
		return
	}

	fmt.Println("Archivo actualizado")
	path2, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path2, "-Tpng", "grafoAVL.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("grafoAVL.png", cmd, os.FileMode(mode))
}

//Recorrido del arbol
func RecorrerArbol(actual **NodoAVL, dot *string) {
	if *actual != nil {
		*dot += "\"" + fmt.Sprint(&(*actual)) + "\"[label = \"" + strconv.Itoa(*&actual.Produ.Codigo) + ""
	}
}

//Reconicimeinto de error para grafos
func ExsisteError() {
	fmt.Println("hola")
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
	ListaTT = nil
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
				indiceT := ((i*len(departamentos)+j)*5 + (k))
				listaAc := Linealizacion[indiceT]
				if listaAc.size != 0 {
					nodG := listaAc.graphNodos(k)
					txtdot += nodG
				}
			}
			rank += "}\n"
			txtdot += rank
			rank = "{rank=same;"

			//Crear conexiones entre nodos
			for k := 0; k < 4; k++ {
				txtdot += "nodo" + strconv.Itoa(k) + "->nodo" + strconv.Itoa(k+1) + "\n"
			}

		}
		txtdot = "digraph G { \nnode[shape=record]\n" + `graph[splines="ortho"]` + "\n"
		rank = "{rank=same;"

		dots := 0
		fmt.Println(txtdot)
		err := ioutil.WriteFile("Grafo"+strconv.Itoa(dots+1)+".dot", []byte(txtdot), 0644)
		if err != nil {
			log.Fatal(err)
		}
		ruta, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(ruta, "-Tpng", "./Tiendas"+strconv.Itoa(dots+1)+".dot").Output()
		mode := int(0777)
		ioutil.WriteFile("Grafo"+strconv.Itoa(dots+1)+".png", cmd, os.FileMode(mode))
	}

}

func (l *ListaD) graphNodos(n int) string {
	nodoA := l.inicio
	txt := ""
	uni := "nodo" + strconv.Itoa(n) + "->"

	//crear nodos de las tienda en lista
	if nodoA != nil {
		for nodoA != nil {
			txt += nodoA.tienda.Nombre + `[label="` + nodoA.tienda.Nombre + "|" + nodoA.tienda.Contacto + "|" + strconv.Itoa(nodoA.tienda.Calificacion) + `"]` + "\n"
			if nodoA.siguiente != nil {
				uni += nodoA.tienda.Nombre + "->" + nodoA.siguiente.tienda.Nombre
			}
			if l.inicio == l.fin {
				uni += nodoA.tienda.Nombre
			}
			nodoA = nodoA.siguiente
		}
		txt += uni + "\n"
		return txt
	}
	return txt
}

func indexR(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor en funcionamiento")
}

func eliminarTienda(w http.ResponseWriter, r *http.Request) {
	var paramDel DeleteShop
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Inserte datos validos")
	}

	json.Unmarshal(reqBody, &paramDel)

	val := len(Linealizacion)
	fmt.Println(val)
	for i := 0; i < val; i++ {
		nodo := Linealizacion[i].inicio
		fmt.Println(i)
		if nodo != nil {
			fmt.Println("Entro a listaD")
			td := Linealizacion[i]
			fmt.Println("pasa linea")

			fmt.Println("casi llega delete")
			//Eliminar nodo
			EliminarNodo(paramDel.Nombre, &td)
			fmt.Println("pasa delete")

			//Salida
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(paramDel)
		}
	}

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
	router.HandleFunc("/Eliminar", eliminarTienda).Methods("POST")
	//router.HandleFunc("/getArreglo", grafo).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
