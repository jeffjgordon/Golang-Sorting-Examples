package main

import (
   "fmt"
   "time"
   "math/rand"
)

func swap(a *int, b *int) {
  t := *a
  *a = *b
  *b = t
}

func track(msg string) (string, time.Time) {
  return msg, time.Now()
}

func duration(msg string, start time.Time) {
  fmt.Printf("%v: %v\n", msg, time.Since(start))
}

func SelectionSort(arr []int) {
  defer duration(track("SelectionSort"))
  var  low int
  n := len(arr)

  for i := 0; i < n; i ++ {
    low = i

    for j := i+1; j < n; j++ {
      if (arr[j] < arr[low]) {
        low = j
      }
    }

    if (arr[i] > arr[low]) {
      swap(&arr[i], &arr[low])
    }
  }

}

func BubbleSort(arr []int) {
  defer duration(track("BubbleSort"))
  n := len(arr)

  for i:=0; i<n; i++ {
    for j:=0; j<(n-i-1); j++ {
      if ( arr[j] > arr[j+1]) {
        swap(&arr[j], &arr[j+1])
      }
    }
  }
}

func BubbleSort2(arr []int, n int) {
  defer duration(track("BubbleSort2"))

  count := 0

  if n == 1 {
    return
  }

  for i:=0; i<n-1; i++ {
    if (arr[i] > arr[i+1]){
      swap(&arr[i], &arr[i+1])
      count ++
    }
  }

  if (count == 0) {
    return
  }

  BubbleSort2(arr, n-1)
}

func InsertionSort(arr []int) {
  defer duration(track("InsertionSort"))
  n := len(arr)
  var key, j int

  for i:=1; i<n; i++ {
    key = arr[i]
    j = i-1

    for (j >= 0 && arr[j] > key) {
      arr[j+1] = arr[j]
      j = j-1
    }

    arr[j+1] = key
  }
}

func Merge(arr []int, left int, mid int, right int){
  var i,j,k int
  n1 := mid - left + 1
  n2 := right - mid

  L := make([]int, n1)
  R := make([]int, n2)

  for i:=0; i<n1; i++ {
    L[i] = arr[left+i]
  }
  for j:=0; j<n2; j++ {
    R[j] = arr[mid + 1 + j]
  }

  i = 0
  j = 0
  k = left
  for ( i<n1 && j < n2) {
    if(L[i] <= R[j]) {
      arr[k] = L[i]
      i++
    } else {
      arr[k] = R[j]
      j++
    }
    k++
  }

  for (i < n1) {
    arr[k] = L[i]
    i++
    k++
  }
  
  for (j < n2) {
    arr[k] = R[j]
    j++
    k++
  }

}

func MergeSort(arr []int, start int, end int) {
  defer duration(track("MergeSort"))

  if (start >= end) {
    return
  }

  mid := start + (end - start) / 2
  MergeSort(arr, start, mid)
  MergeSort(arr, mid+1, end)

  Merge(arr, start,mid,end)

}

func partition (arr []int, low int, high int) int {
  pivot := arr[high]
  i := (low - 1)

  for j:=low; j<=high-1; j++ {
    if(arr[j] < pivot) {
      i++
      swap(&arr[i], &arr[j])
    }
  }
  swap(&arr[i+1], &arr[high])
  return (i+1)
}

func QuickSort(arr []int, low int, high int) {
  defer duration(track("QuickSort"))
  
  if (low < high) {
    pi := partition(arr, low, high)
    
    QuickSort(arr,low,pi-1)
    QuickSort(arr,pi+1,high)
  }
}

func HeapSort(arr []int) {

}

func CountingSort(arr []int) {

}

func RadixSort(arr []int) {

}

func BucketSort(arr []int) {

}

func generateRandomSlice(size int) []int {
  slice:= make([]int, size, size)
  rand.Seed(time.Now().UnixNano())
  for i:=0; i< size; i++ {
    slice[i] = rand.Intn(999) - rand.Intn(999)
  }
  return slice
}

func main () {
  defer duration(track("Main"))
  /*
  arr1 := [3]int{4, 5, 6}
  arr2 := [2]int{1,2}
  arr3 := [1]int{3}
  var sortableSlice []int
  */
  arrSize := 5
/*
  sortableSlice = append(sortableSlice, arr1[:]...)
  sortableSlice = append(sortableSlice, arr2[:]...)
  sortableSlice = append(sortableSlice, arr3[:]...)
*/


  selectionSlice := generateRandomSlice(arrSize)
  fmt.Println("Selection Sort \n---\nPre-Sorted:", selectionSlice)
  SelectionSort(selectionSlice)
  fmt.Println("After Sort:",selectionSlice, "\n\n========\n")

  bubbleSlice1 := generateRandomSlice(arrSize)
  fmt.Println("Standard Bubble Sort \n---\nPre-Sorted:", bubbleSlice1)
  BubbleSort(bubbleSlice1)
  fmt.Println("After Sort:",bubbleSlice1,"\n\n========\n")

  bubbleSlice2 := generateRandomSlice(arrSize)
  fmt.Println("Recursive Bubble Sort \n---\nPre-Sorted:", bubbleSlice2)
  BubbleSort2(bubbleSlice2, len(bubbleSlice2)-1)
  fmt.Println("After Sort:",bubbleSlice2,"\n\n========\n")

  insertionSlice := generateRandomSlice(arrSize)
  fmt.Println("Insertion Sort \n---\nPre-Sorted:", insertionSlice)
  InsertionSort(insertionSlice)
  fmt.Println("After Sort:",insertionSlice,"\n========\n")

  mergeSlice := generateRandomSlice(arrSize)
  fmt.Println("Merge Sort \n---\nPre-Sorted:", mergeSlice)
  MergeSort(mergeSlice,0,len(mergeSlice)-1)
  fmt.Println("After Sort:",mergeSlice,"\n========\n")

  quickSlice := generateRandomSlice(arrSize)
  fmt.Println("Quick Sort \n---\nPre-Sorted:",quickSlice)
  QuickSort(quickSlice, 0, len(quickSlice)-1)
  fmt.Println("After Sort:",quickSlice,"\n========\n")

//  fmt.Println("Pre-sorted list:",sortableSlice)
//  SelectionSort(sortableSlice)
//  BubbleSort(sortableSlice)
//  BubbleSort2(sortableSlice,len(sortableSlice))
//  InsertionSort(sortableSlice)
//  MergeSort(sortableSlice, 0, len(sortableSlice)-1)
//  QuickSort(sortableSlice,0, len(sortableSlice)-1)
//  fmt.Println("Post-sorted list:",sortableSlice)



}