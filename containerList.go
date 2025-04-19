// Container List adalah package yang berisi fungsi-fungsi untuk mengelola list
// biasanya list yg double linked list apa itu ? yaitu list yang memiliki dua pointer
// satu ke next dan satu lagi ke prev
// untuk apa? biasanya dipakai untuk queue/antrian, stack/tumpukan, dll

package main

import "container/list"

func main(){
	data := list.New()
	data.PushBack(1)
	data.PushBack(2)
	data.PushBack(3)

	// data.Front() adalah pointer ke elemen pertama
	// data.Back() adalah pointer ke elemen terakhir
	// data.Next() adalah pointer ke elemen selanjutnya
	for e := data.Front(); e != nil; e = e.Next() {
		println(e.Value.(int))
	}
}