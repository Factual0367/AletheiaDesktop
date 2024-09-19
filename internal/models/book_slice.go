package models

type BookSlice []*Book

func (bs BookSlice) Len() int           { return len(bs) }
func (bs BookSlice) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
func (bs BookSlice) Less(i, j int) bool { return bs[i].Title < bs[j].Title }
