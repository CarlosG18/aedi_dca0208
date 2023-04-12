package list

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	cabeca *No
	tam    int
}

type No struct {
	value int
	prox  *No
}

func (linkedlist *LinkedList) Init() {
	linkedlist.cabeca = &No{}
	linkedlist.tam = 0
}

func (linkedlist *LinkedList) Add(value int) {
	var aux *No = linkedlist.cabeca

	for aux.prox != nil {
		aux = aux.prox
	}
	if aux.prox == nil {
		aux.value = value
		aux.prox = &No{}
	}
	linkedlist.tam++
}

func (linkedlist *LinkedList) Remove() error {
	var aux *No = linkedlist.cabeca
	var auxAnt *No

	if linkedlist.tam > 0 {
		for aux.prox != nil {
			auxAnt = aux
			aux = aux.prox
		}
		if aux.prox == nil {
			auxAnt.prox = nil
		}
		linkedlist.tam--
		return nil
	} else {
		return errors.New("Linkedlist não possui elementos!")
	}
}

func (linkedlist *LinkedList) RemoveOnIndex(index int) error {
	cont := 0
	var aux *No = linkedlist.cabeca
	var aux1 *No = aux.prox

	if linkedlist.tam > 0 && (index < linkedlist.tam && index >= 0) {
		if index == 0 {
			linkedlist.cabeca = aux1
		} else {
			for cont < index-1 {
				aux = aux.prox
				aux1 = aux.prox
				cont++
			}
			aux.prox = aux1.prox
		}
		linkedlist.tam--
		return nil
	} else {
		return errors.New("error: Não foi possível acessar o index!")
	}
}

func (linkedlist *LinkedList) AddOnIndex(value int, index int) error {
	cont := 0
	var aux *No = linkedlist.cabeca
	var aux1 *No = aux.prox
	var new_num *No = &No{}

	if index < 0 || index > linkedlist.tam {
		return errors.New("error: não foi possivel acessar o index!")
	} else {
		if index == 0 {
			linkedlist.cabeca = new_num
			new_num.value = value
			new_num.prox = aux
		} else {
			for cont < index-1 {
				aux = aux.prox
				aux1 = aux.prox
				cont++
			}
			aux.prox = new_num
			new_num.value = value
			new_num.prox = aux1
		}
		linkedlist.tam++
		return nil
	}
}

func (linkedlist *LinkedList) Size() int {
	return linkedlist.tam
}

func (linkedlist *LinkedList) Get(index int) (int, error) {
	cont := 0
	var aux *No = linkedlist.cabeca

	if index < 0 || index > linkedlist.tam {
		return -1, errors.New("error: não foi possivel acessar o index!")
	} else {
		for aux.prox != nil && cont != index {
			aux = aux.prox
			cont++
		}
		return aux.value, nil
	}
}

func (linkedlist *LinkedList) Set(value int, index int) error {
	cont := 0
	var aux *No = linkedlist.cabeca

	if index < 0 || index >= linkedlist.tam {
		return errors.New("error: não foi possivel acessar o index!")
	} else {
		for aux.prox != nil && cont != index {
			aux = aux.prox
			cont++
		}
		aux.value = value
		return nil
	}
}

func (linkedlist *LinkedList) Invert() {
	//anterior := No{}
	atual := linkedlist.cabeca
	prox := atual.prox
	var ant *No = nil
	fmt.Println()
	fmt.Println()

	for atual != nil {
		fmt.Println("anterior = ", ant, " atual = ", atual, " proximo = ", prox)
		if atual.prox != nil {
			prox = atual.prox
		}
		ant = atual
		atual = atual.prox

		prox = atual.prox
		atual.prox = ant
		ant = atual
		atual = prox
	}
	//linkedlist.cabeca = atual
}

func (linkedlist *LinkedList) Print() {
	aux := linkedlist.cabeca
	for aux != nil {
		fmt.Print(aux.value)
		if aux.prox != nil {
			fmt.Print(",")
		}
		aux = aux.prox
	}
}
