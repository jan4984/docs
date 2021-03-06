## engine
* interpreter

x86 CPU是x86机器语言的解释器
* compiler

一种语言转化成另一种语言
* AOT(ahead of time) JIT (just in time)

运行前转换，运行时转换。第一次运行特别慢？

http://softwareengineering.stackexchange.com/questions/246094/understanding-the-differences-traditional-interpreter-jit-compiler-jit-interp/269878#269878

## 运行环境
* browser

HTML/CSS解析渲染；其他扩展；菜单/收藏等GUI；js引擎；js语言访问HTML等扩展的API实现
![H5相关标准](https://upload.wikimedia.org/wikipedia/commons/thumb/7/7f/HTML5_APIs_and_related_technologies_taxonomy_and_status.svg/1133px-HTML5_APIs_and_related_technologies_taxonomy_and_status.svg.png)

* nodejs

js引擎；js访问OS资源（内存、文件系统、网络）的API实现

* Nashorn(JVM) IronJS(.net) v8(c++) JavaScriptCore(iOS-ObjectiveC&Swift) ...

js引擎；各种语言的互交互（函数调用、数据类型转换、包装）

## 预编译

* Babel Webpack

## JavaScript
* primitives

```javascript
var a=1.2;
var a="a";
var a='a';
var a=`hello"fd
fd

s'123sdf'`;
function a(){}
var a = function(p){}
var a=v=>5；
var c=(a,b)=>{console.log(a);return b;};
var a=NaN;
var a=undefined;
var a=null;
```
* closures
const var let

https://jsfiddle.net/v7gjv/?utm_source=website&utm_medium=embed&utm_campaign=v7gjv

* && ||

```javascript
const a = a||{v:'init value'};
const b = a && a.v;
```
* array & object

a[] a.b a['b'] delete(a.b)

* export class extends super new constructor() bind apply arguments

* destructuring spread

```javascript
let a=5; const b={a};//b={a:5}
{a}=b;//a=b.a
const c={...b,v:'dd'};//c={a:5,v:'dd'}
const d=[1,2,3];
const e=[...d,4,5];//e=[1,2,3,4,5]
```

* concat join pop push slice sort filter shift unshift splice map forEach find findIndex lastIndexOf reduce 

https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array

* json

```javascript
JSON.stringify({a:5});//{"a":5}
cosnt v = JSON.parse('{"a":5}');//v={a:5}
```

* Promise

保持异步调用时的业务代码组织
```javascript
function asyncGet(url){
    return new Promise((resolve, reject)=> {
        var xhr = new XMLHttpRequest();
        xhr.open("GET", url, true);
        xhr.onload = ()=>{
            if (xhr.readyState === 4) {
                if (xhr.status === 200 || xhr.status === 0) {//status 0 for local file
                    resolve(xhr.responseText);
                } else {
                    console.log('get', url, 'failed:', xhr.statusText);
                }
            }
        };
        xhr.onerror = function (e) {
            console.log('get', url, 'failed:', xhr.statusText);
            reject(xhr.statusText);
        };
        xhr.send(null);
    });
}
asyncGet(path).then(res=>this.loadJson(res));
```

https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise


## Why JavaScript

* corss platforms with **GUI** by Browser
* live update by Web
* more dynamic with interpreter at runtime
* more and more capbilities by Browser(3D/Audio/Video/Location/VR/service worker/...)
* more and more frameworks & tools by Companies(Google/MS/FB/Alibaba/... & start-ups)
