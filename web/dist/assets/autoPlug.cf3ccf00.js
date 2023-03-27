/*! 
 Build based on gin-vue-admin 
 Time : 1679909256000 */
import{a as e,o as l,c as a,d as o,e as u,w as s,a0 as n,U as t,aC as d,J as p,d4 as r,d5 as i,f as m,h as c,j as b,bm as f,b as V,E as y}from"./index.61dd0a2d.js";import{E as k,a as h}from"./form-item.a3ced827.js";/* empty css               *//* empty css              */import"./tag.3ecd1fed.js";import{E as g,a as v}from"./select.31449efc.js";import"./scrollbar.ef9dc6d6.js";/* empty css               */import{E as U}from"./checkbox.384d52e9.js";import{b as _}from"./stringFun.1a90c7e7.js";import{k as j}from"./autoCode.3f49bdae.js";import{_ as R}from"./_plugin-vue_export-helper.cdc0426e.js";import"./_Uint8Array.0caf8bf8.js";import"./_initCloneObject.e29c0e59.js";import"./index.cbcb4c8f.js";import"./index.ba9b69da.js";import"./strings.ed5ba909.js";import"./isEqual.3dfeee36.js";const q={class:"gva-table-box"},C=R({__name:"autoPlug",setup(R){const C=e({plugName:"",routerGroup:"",hasGlobal:!0,hasRequest:!0,hasResponse:!0,global:[{key:"",type:"",desc:""}],request:[{key:"",type:"",desc:""}],response:[{key:"",type:"",desc:""}]}),G=()=>{C.plugName=_(C.plugName)},x=async()=>{if(!C.plugName||!C.routerGroup)return void b.error("插件名称和插件路由组为必填项");if(C.hasGlobal){if(C.global.some((e=>{if(!e.key||!e.type)return!0})))return void b.error("全局属性的key和type为必填项")}if(C.hasRequest){if(C.request.some((e=>{if(!e.key||!e.type)return!0})))return void b.error("请求属性的key和type为必填项")}if(C.hasResponse){if(C.response.some((e=>{if(!e.key||!e.type)return!0})))return void b.error("响应属性的key和type为必填项")}0===(await j(C)).code&&f("创建成功，插件已自动写入后端plugin目录下，请按照自己的逻辑进行创造")},w=e=>{e.push({key:"",value:""})},N=(e,l)=>{1!==e.length?e.splice(l,1):b.warning("至少有一个全局属性")};return(e,b)=>{const f=V,_=k,j=U,R=g,E=v,A=y,B=h;return l(),a("div",null,[o("div",q,[u(B,{"label-width":"140px",class:"plug-form"},{default:s((()=>[u(_,{label:"插件名"},{default:s((()=>[u(f,{modelValue:C.plugName,"onUpdate:modelValue":b[0]||(b[0]=e=>C.plugName=e),placeholder:"必填（英文大写字母开头）",onBlur:G},null,8,["modelValue"])])),_:1}),u(_,{label:"路由组"},{default:s((()=>[u(f,{modelValue:C.routerGroup,"onUpdate:modelValue":b[1]||(b[1]=e=>C.routerGroup=e),placeholder:"将会作为插件路由组使用"},null,8,["modelValue"])])),_:1}),u(_,{label:"使用全局属性"},{default:s((()=>[u(j,{modelValue:C.hasGlobal,"onUpdate:modelValue":b[2]||(b[2]=e=>C.hasGlobal=e)},null,8,["modelValue"])])),_:1}),C.hasGlobal?(l(),n(_,{key:0,label:"全局属性"},{default:s((()=>[(l(!0),a(t,null,d(C.global,((e,n)=>(l(),a("div",{key:n,class:"plug-row"},[o("span",null,[u(f,{modelValue:e.key,"onUpdate:modelValue":l=>e.key=l,placeholder:"key 必填"},null,8,["modelValue","onUpdate:modelValue"])]),o("span",null,[u(E,{modelValue:e.type,"onUpdate:modelValue":l=>e.type=l,placeholder:"type 必填"},{default:s((()=>[u(R,{label:"string",value:"string"}),u(R,{label:"int",value:"int"}),u(R,{label:"float32",value:"float32"}),u(R,{label:"float64",value:"float64"}),u(R,{label:"bool",value:"bool"})])),_:2},1032,["modelValue","onUpdate:modelValue"])]),o("span",null,[u(f,{modelValue:e.desc,"onUpdate:modelValue":l=>e.desc=l,placeholder:"备注 必填"},null,8,["modelValue","onUpdate:modelValue"])]),o("span",null,[u(A,{icon:p(r),circle:"",onClick:b[3]||(b[3]=e=>w(C.global))},null,8,["icon"])]),o("span",null,[u(A,{icon:p(i),circle:"",onClick:e=>N(C.global,n)},null,8,["icon","onClick"])])])))),128))])),_:1})):m("",!0),u(_,{label:"使用Request"},{default:s((()=>[u(j,{modelValue:C.hasRequest,"onUpdate:modelValue":b[4]||(b[4]=e=>C.hasRequest=e)},null,8,["modelValue"])])),_:1}),C.hasRequest?(l(),n(_,{key:1,label:"Request"},{default:s((()=>[(l(!0),a(t,null,d(C.request,((e,n)=>(l(),a("div",{key:n,class:"plug-row"},[o("span",null,[u(f,{modelValue:e.key,"onUpdate:modelValue":l=>e.key=l,placeholder:"key 必填"},null,8,["modelValue","onUpdate:modelValue"])]),o("span",null,[u(E,{modelValue:e.type,"onUpdate:modelValue":l=>e.type=l,placeholder:"type 必填"},{default:s((()=>[u(R,{label:"string",value:"string"}),u(R,{label:"int",value:"int"}),u(R,{label:"float32",value:"float32"}),u(R,{label:"float64",value:"float64"}),u(R,{label:"bool",value:"bool"})])),_:2},1032,["modelValue","onUpdate:modelValue"])]),o("span",null,[u(f,{modelValue:e.desc,"onUpdate:modelValue":l=>e.desc=l,placeholder:"备注 必填"},null,8,["modelValue","onUpdate:modelValue"])]),o("span",null,[u(A,{icon:p(r),circle:"",onClick:b[5]||(b[5]=e=>w(C.request))},null,8,["icon"])]),o("span",null,[u(A,{icon:p(i),circle:"",onClick:e=>N(C.request,n)},null,8,["icon","onClick"])])])))),128))])),_:1})):m("",!0),u(_,{label:"使用Response"},{default:s((()=>[u(j,{modelValue:C.hasResponse,"onUpdate:modelValue":b[6]||(b[6]=e=>C.hasResponse=e)},null,8,["modelValue"])])),_:1}),C.hasResponse?(l(),n(_,{key:2,label:"Response"},{default:s((()=>[(l(!0),a(t,null,d(C.response,((e,n)=>(l(),a("div",{key:n,class:"plug-row"},[o("span",null,[u(f,{modelValue:e.key,"onUpdate:modelValue":l=>e.key=l,placeholder:"key 必填"},null,8,["modelValue","onUpdate:modelValue"])]),o("span",null,[u(E,{modelValue:e.type,"onUpdate:modelValue":l=>e.type=l,placeholder:"type 必填"},{default:s((()=>[u(R,{label:"string",value:"string"}),u(R,{label:"int",value:"int"}),u(R,{label:"float32",value:"float32"}),u(R,{label:"float64",value:"float64"}),u(R,{label:"bool",value:"bool"})])),_:2},1032,["modelValue","onUpdate:modelValue"])]),o("span",null,[u(f,{modelValue:e.desc,"onUpdate:modelValue":l=>e.desc=l,placeholder:"备注 必填"},null,8,["modelValue","onUpdate:modelValue"])]),o("span",null,[u(A,{icon:p(r),circle:"",onClick:b[7]||(b[7]=e=>w(C.response))},null,8,["icon"])]),o("span",null,[u(A,{icon:p(i),circle:"",onClick:e=>N(C.response,n)},null,8,["icon","onClick"])])])))),128))])),_:1})):m("",!0),u(_,null,{default:s((()=>[u(A,{type:"primary",onClick:x},{default:s((()=>[c("创建")])),_:1})])),_:1})])),_:1})])])}}},[["__scopeId","data-v-6641fd51"]]);export{C as default};
