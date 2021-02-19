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

type listaT []Tienda

var matriz [][][]listaT
var indice [27]string
var departamentos [][]string
var vector []listaT

func fullM() {
	cajita = make([][][]Lista_doble, len(a.Datos))
	fmt.Println(len(a.Datos))
	for i := 0; i < len(cajita); i++ {
		cajita[i] = make([][]Lista_doble, len(a.Datos[0].Departamentos))
		for j := 0; j < len(cajita[i]); j++ {
			cajita[i][j] = make([]Lista_doble, 5)
		}
	}
}

var ListaTiendas = listaT{
	{
		Nombre:       "iShop",
		Descripcion:  "Expertos en Apple",
		Contacto:     "5529-0756",
		Calificacion: 5,
	},
}

func getTiendas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ListaTiendas)
}

func crearTienda(w http.ResponseWriter, r *http.Request) {
	var newtienda Tienda
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Inserte datos validos")
	}

	json.Unmarshal(reqBody, &newtienda)

	ListaTiendas = append(ListaTiendas, newtienda)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newtienda)
}

func searchtienda(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	nombretienda := vars["nombre"]

	for i, tienda := range ListaTiendas {
		if tienda.Nombre == nombretienda {

		}
	}

}

func deleteTienda(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tiendaNombre := vars["nombre"]

	for _, tienda := range ListaTiendas {
		if tienda.Nombre == tiendaNombre {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tienda)
		}
	}
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
