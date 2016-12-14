https://github.com/tastejs/todomvc/tree/gh-pages/examples/react

###Component Element

**element** 是一个 VirtualDOM 中的节点，包含类型、属性和子 **element**。类型可以是字符串类型或者 **component** 类型。如果是字符串类型 X(如span/div等) 则最终会转化为 HTML 元素 X；如果是 **component** 类型则会从 `render()` 方法进一步扩展。需要注意的，属性可以是函数，因为 JavaScript 中函数是一等公民。

**component** 是继承 `React.Component` 的 JavaScript 类，必须有 `render()` 方法和 props 属性，可以有 state 属性，其他方面就是一个正常的 JavaScript 类。`render()` 方法用于返回一个 **element**；props 和 state 用于控制返回的 **element**。props 是由父元素指定的，在 **component** 内不可变；state 是自己的状态，可以在`render()`之外变化，所以如果不需要维护自己的长期状态，就可以不用 state。我们也可以通过 props 和 state 之外的数据创建 **element**，但是这样就脱离了 React 框架设计的生命周期控制，也就失去了使用 React 框架的意义。

一个 **component** 类型的 **element** 就会对应一个 **component** 类的实例，但两者并不是一个东西。一个仅仅是Virtual DOM中的节点，一个是JavaScript类的实例化对象。通过 **element** 的 ref 属性，可以找到对应的 **component** 实例。

我们应该尽量选择不带内部状态的 **component**，可变量越少，程序越简单易维护。 

https://gist.github.com/sebmarkbage/fcb1b6ab493b0c77d589

###Virtual DOM

Virtual DOM 中的所有 Element 最终会被映射到浏览器 DOM 中， 达到GUI效果。

可以在 JavaScript 中持有并传递 ReactElement，但如果需要插入到 Virtual DOM 则需要在 `render()` 方法中返回。

从 Virtual DOM 对应到 DOM 可以用 `ReactDOM.findDOMNode(comp)`；从 DOM 对应到 Virtual DOM 可以用 event。比如 `<button onClick={()=>console.log(this.name + ' clicked')}/>`

http://facebook.github.io/react/docs/reconciliation.htm

### Component Instance Lifecycle

http://javascript.tutorialhorizon.com/2014/09/13/execution-sequence-of-a-react-components-lifecycle-methods/
