package bubbleSort

//Request of RPC call
type Args struct {
	Array []int
}

//Server service for RPC
type BubbleSort int

//Result of RPC call
type Result []int

//Sort RPC function 
func (t *BubbleSort) Sort(args Args, Result *[]int) error {
	arr := args.Array
	n := len(arr)

    //Implementation of Bubble Sort algorithm
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                //Swap arr[j] and arr[j+1]
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }

	*Result =  arr 
	return nil
}