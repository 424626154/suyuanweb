package models



type Page struct {
    PageNo     int
    PageSize   int
    TotalPage  int
    TotalCount int
    FirstPage  bool
    LastPage   bool
    List       interface{}
}