package array

type MyArray struct {
	Data []interface{}
	Size int
}

func NewMyArray(size int) *MyArray {
	if size == 0 {
		return nil
	}

	return &MyArray{
		make([]interface{}, size),
		size,
	}
}

// 增
func (m *MyArray) AddItem(item interface{}) {
	if len(m.Data) == m.Size {
		m.resize()
	}
	m.Data[m.Size] = item
	m.Size++
}

func (m *MyArray) AddItems(items ...interface{}) {
	for item := range items {
		m.AddItem(item)
	}
}

func (m *MyArray) resize() {
	newArray := make([]interface{}, 2*m.Size)
	for i := 0; i <= m.Size; i++ {
		newArray[i] = m.Data[i]
	}
	m.Data = newArray
}

// 查
func (m *MyArray) FindItem(target interface{}) int {
	for i := 0; i < m.Size; i++ {
		if m.Data[i] == target {
			return i
		}
	}
	return -1
}

func (m *MyArray) FindAllItems(target interface{}) (result []int) {
	for i := 0; i < m.Size; i++ {
		if m.Data[i] == target {
			result = append(result, i)
		}
	}
	return
}

func (m *MyArray) GetItem(index int) interface{} {
	if index > len(m.Data) || index < 0 {
		return nil
	}
	return m.Data[index]
}

func (m *MyArray) GetItems(indexes ...int) []interface{} {
	result := make([]interface{}, len(indexes))
	for i := range indexes {
		if i >= 0 && i < m.Size {
			result = append(result, m.Data[i])
		}
	}

	return result
}

func (m *MyArray) Contains(target interface{}) bool {
	if m.FindItem(target) == -1 {
		return false
	}
	return true
}

// 删
func (m *MyArray) Remove(target interface{}) {
	newArray := make([]interface{}, len(m.Data))
	for i := 0; i < m.Size; i++ {
		if m.Data[i] != target {
			newArray = append(newArray, m.Data[i])
		}
	}
}

func (m *MyArray) RemoveWithIndex(index int) {
	if index < 0 || index > len(m.Data) {
		panic("Illegal index.")
	}
	newArray := make([]interface{}, len(m.Data))
	for i := 0; i < m.Size; i++ {
		if i != index {
			newArray[i] = m.Data[i]
		}
	}
	m.Data = newArray
}

// 改
func (m *MyArray) Set(index int, item interface{}) {
	if index < 0 || index > len(m.Data) {
		panic("Illegal index.")
	}
	m.Data[index] = item
}

func (m *MyArray) GetCapacity() int {
	return len(m.Data)
}

func (m *MyArray) GetSize() int {
	return m.Size
}

func (m *MyArray) IsEmpty() bool {
	return m.Size == 0
}
