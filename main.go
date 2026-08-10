package main
import ("fmt";"reflect";"strings")
const testSuite = "cron-tick-983464"
type TestCase struct{Name string;Input interface{};Expected interface{}}
func assertEqual(name string, got, want interface{}) bool{ok:=reflect.DeepEqual(got,want);status:="PASS";if !ok{status="FAIL"};fmt.Printf("[%s] %s: %s (got=%v want=%v)\n",testSuite,status,name,got,want);return ok}
func runSuite(cases []TestCase, fn func(interface{}) interface{}){pass,fail:=0,0;for _,tc:=range cases{if assertEqual(tc.Name,fn(tc.Input),tc.Expected){pass++}else{fail++}};fmt.Printf("[%s] Results: %d passed, %d failed\n",testSuite,pass,fail)}
func main(){cases:=[]TestCase{{"upper","hello","HELLO"},{"trim","  hi  ","hi"},{"replace","foo-bar","foo_bar"}};runSuite(cases,func(in interface{}) interface{}{s:=in.(string);s=strings.TrimSpace(s);s=strings.ToUpper(s);s=strings.ReplaceAll(s,"-","_");return s})}
