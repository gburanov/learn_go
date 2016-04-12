package main

type ListElement struct {
    Value int
    Next *ListElement
}

func removeDuplicates(start *ListElement) {
    next_element := *start.Next

    if next_element == nil {
        return
    }
    next_val := next_element.Value
    if Value == next_val {
        new_next = *next_element.Next
        if (new_next == nil) {
            next_val.Next = nil
            return
        }
        next_val.Next = new_next
        removeDuplicates(start)
        return


    }
    removeDuplicates(*start.Next)
}

func removeDuplicates(start *ListElement) {
    if start == nil {
        return
    }

    for elem := start; elem != nil;  {
        next_element := *elem.Next
        if next_element == nil {
            break
        }
        next_val := next_element.Value
        if Value == next_val {
            new_next = *next_element.Next
            if (new_next == nil) {
                next_val.Next = nil
                break
            }
            next_val.Next = new_next
        } else {
            elem = *elem.Next
        }
    }

}
