package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Lista Doble para alacenar tiendas------------------------------------
//Nodos
type Nodo struct {
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

type Datos struct {
	Seccion []Seccion `json:Datos`
}

type Seccion struct {
	Indice        string         `json:Indice`
	Departamentos []Departamento `json:Departamento`
}

type Departamento struct {
	Nombre  string   `json:Nombre`
	Tiendas []Tienda `json:Tienda`
}

type Tienda struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	Contacto     string `json:Contacto`
	Calificacion int    `json:Calificacion`
}

var indice [27]string
var departamentos []string
var linealizacion []ListaD

func matriz(a *Datos) {

	departamentos := make([]string, len(a.Seccion[0].Departamentos))

	for i := 0; i < len(a.Seccion); i++ {
		indice[i] = a.Seccion[i].Indice
	}

	for i := 0; i < len(a.Seccion[0].Departamentos); i++ {
		departamentos[i] = a.Seccion[i].Departamentos[i].Nombre
	}

	//Estableciendo tamaño al vector de la matrzi linealizada
	linealizacion = make([]ListaD, (len(a.Seccion) * len(a.Seccion) * len(a.Seccion[0].Departamentos) * 5))

}

func getTiendas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(ListaTiendas)
}

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

func llenarMatriz(da *Datos) {
	for i := 0; i < len(da.Seccion); i++ {
		for j := 0; j < len(da.Seccion[i].Departamentos); j++ {
			for k := 0; k < len(da.Seccion[i].Departamentos[j].Tiendas); k++ {

			}
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

func indexR(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor en funcionamiento")
}

func main() {
	fmt.Print("Hola")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexR)
	router.HandleFunc("/tiendas", getTiendas).Methods("GET")
	router.HandleFunc("/tiendas", crearTienda).Methods("POST")
	router.HandleFunc("/tiendas/{nombre}", searchtienda).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
