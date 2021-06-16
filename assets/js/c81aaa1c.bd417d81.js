(self.webpackChunkgolang=self.webpackChunkgolang||[]).push([[190],{3905:function(e,t,r){"use strict";r.d(t,{Zo:function(){return s},kt:function(){return g}});var n=r(7294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function o(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?o(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var c=n.createContext({}),p=function(e){var t=n.useContext(c),r=t;return e&&(r="function"==typeof e?e(t):i(i({},t),e)),r},s=function(e){var t=p(e.components);return n.createElement(c.Provider,{value:t},e.children)},m={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},u=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,o=e.originalType,c=e.parentName,s=l(e,["components","mdxType","originalType","parentName"]),u=p(r),g=a,d=u["".concat(c,".").concat(g)]||u[g]||m[g]||o;return r?n.createElement(d,i(i({ref:t},s),{},{components:r})):n.createElement(d,i({ref:t},s))}));function g(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=r.length,i=new Array(o);i[0]=u;var l={};for(var c in t)hasOwnProperty.call(t,c)&&(l[c]=t[c]);l.originalType=e,l.mdxType="string"==typeof e?e:a,i[1]=l;for(var p=2;p<o;p++)i[p]=r[p];return n.createElement.apply(null,i)}return n.createElement.apply(null,r)}u.displayName="MDXCreateElement"},276:function(e,t,r){"use strict";r.r(t),r.d(t,{frontMatter:function(){return l},metadata:function(){return c},toc:function(){return p},default:function(){return m}});var n=r(2122),a=r(9756),o=(r(7294),r(3905)),i=["components"],l={sidebar_position:5},c={unversionedId:"tutorial-extras/conversion",id:"tutorial-extras/conversion",isDocsHomePage:!1,title:"Type Conversion",description:"Name Convention",source:"@site/docs/tutorial-extras/05-conversion.md",sourceDirName:"tutorial-extras",slug:"/tutorial-extras/conversion",permalink:"/docs/tutorial-extras/conversion",editUrl:"https://github.com/facebook/docusaurus/edit/master/website/docs/tutorial-extras/05-conversion.md",version:"current",sidebarPosition:5,frontMatter:{sidebar_position:5},sidebar:"tutorialSidebar",previous:{title:"Reflection Helper",permalink:"/docs/tutorial-extras/reflection-package"}},p=[{value:"Name Convention",id:"name-convention",children:[]},{value:"Declare Syntax",id:"declare-syntax",children:[]}],s={toc:p};function m(e){var t=e.components,r=(0,a.Z)(e,i);return(0,o.kt)("wrapper",(0,n.Z)({},s,r,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h2",{id:"name-convention"},"Name Convention"),(0,o.kt)("ol",null,(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("p",{parentName:"li"},"Variable start with a letter"),(0,o.kt)("ol",{parentName:"li"},(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("p",{parentName:"li"},(0,o.kt)("strong",{parentName:"p"},(0,o.kt)("em",{parentName:"strong"},"With capital letter is exported and used outside of package"))),(0,o.kt)("pre",{parentName:"li"},(0,o.kt)("code",{parentName:"pre",className:"language-go"},' package main\n\n import "fmt"\n\n // Greeting is exported\n var Greeting string = "Hello World"\n\n func main() {\n   // using Greeting variable\n    fmt.Println(Greeting)\n }\n'))),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("p",{parentName:"li"},"With lowercase letter is not exported")))),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("p",{parentName:"li"},"Variable with multiple words"),(0,o.kt)("ol",{parentName:"li"},(0,o.kt)("li",{parentName:"ol"},"Attached together without space"),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("strong",{parentName:"li"},(0,o.kt)("em",{parentName:"strong"},"Example")),(0,o.kt)("pre",{parentName:"li"},(0,o.kt)("code",{parentName:"pre",className:"language-go"},"var alochymInt    int\nvar alochymString string\nvar alochymBool   bool\nvar AlochymUnit   float64\n")))))),(0,o.kt)("h2",{id:"declare-syntax"},"Declare Syntax"),(0,o.kt)("ol",null,(0,o.kt)("li",{parentName:"ol"},"Declare with assigning value",(0,o.kt)("pre",{parentName:"li"},(0,o.kt)("code",{parentName:"pre",className:"language-go"},"var var_name type = value\n"))),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("strong",{parentName:"li"},(0,o.kt)("em",{parentName:"strong"},"Example")),(0,o.kt)("pre",{parentName:"li"},(0,o.kt)("code",{parentName:"pre",className:"language-go"},"var alochym int = 1\n"))),(0,o.kt)("li",{parentName:"ol"},"Declare without assigning value => zero value",(0,o.kt)("pre",{parentName:"li"},(0,o.kt)("code",{parentName:"pre",className:"language-go"},"var var_name type\n"))),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("strong",{parentName:"li"},(0,o.kt)("em",{parentName:"strong"},"Example")),(0,o.kt)("pre",{parentName:"li"},(0,o.kt)("code",{parentName:"pre",className:"language-go"},'var alochym       int     => zero value = 0\nvar alochymFloat  float64 => zero value = 0.0\nvar alochymString string  => zero value = ""\nvar alochymBool   bool    => zero value = false\n'))),(0,o.kt)("li",{parentName:"ol"},"Using variable without declaring",(0,o.kt)("pre",{parentName:"li"},(0,o.kt)("code",{parentName:"pre",className:"language-go"},"alochym := 1\n")))))}m.isMDXComponent=!0}}]);