/*! 
 Build based on gin-vue-admin 
 Time : 1680141068000 */
import{c1 as e,u as a,r as s,bf as l,z as t,D as n,ck as u,bH as o,o as r,c as i,e as v,w as m,U as c,aC as d,a0 as p,d as y,a2 as f,J as g,h,a3 as b,cl as S,ah as I,af as q,ag as O}from"./index.9fcd1cf4.js";import{E as k,a as w}from"./tab-pane.82f4a7bd.js";import"./strings.bddfd219.js";const C={class:"router-history"},N=["tab"],J={name:"HistoryComponent"},x=Object.assign(J,{setup(J){const x=e(),j=a(),E=e=>e.name+JSON.stringify(e.query)+JSON.stringify(e.params),V=s([]),A=s(""),P=s(!1),T=l(),L=e=>e.name+JSON.stringify(e.query)+JSON.stringify(e.params),R=s(0),H=s(0),U=s(!1),_=s(!1),z=s(""),D=t((()=>T.userInfo.authority.defaultRouter)),K=()=>{V.value=[{name:D.value,meta:{title:"首页"},query:{},params:{}}],j.push({name:D.value}),P.value=!1,sessionStorage.setItem("historys",JSON.stringify(V.value))},X=()=>{let e;const a=V.value.findIndex((a=>(E(a)===z.value&&(e=a),E(a)===z.value))),s=V.value.findIndex((e=>E(e)===A.value));V.value.splice(0,a),a>s&&j.push(e),sessionStorage.setItem("historys",JSON.stringify(V.value))},Y=()=>{let e;const a=V.value.findIndex((a=>(E(a)===z.value&&(e=a),E(a)===z.value))),s=V.value.findIndex((e=>E(e)===A.value));V.value.splice(a+1,V.value.length),a<s&&j.push(e),sessionStorage.setItem("historys",JSON.stringify(V.value))},$=()=>{let e;V.value=V.value.filter((a=>(E(a)===z.value&&(e=a),E(a)===z.value))),j.push(e),sessionStorage.setItem("historys",JSON.stringify(V.value))},B=e=>{if(!V.value.some((a=>((e,a)=>{if(e.name!==a.name)return!1;if(Object.keys(e.query).length!==Object.keys(a.query).length||Object.keys(e.params).length!==Object.keys(a.params).length)return!1;for(const s in e.query)if(e.query[s]!==a.query[s])return!1;for(const s in e.params)if(e.params[s]!==a.params[s])return!1;return!0})(a,e)))){const a={};a.name=e.name,a.meta={...e.meta},delete a.meta.matched,a.query=e.query,a.params=e.params,V.value.push(a)}window.sessionStorage.setItem("activeValue",E(e))},F=s({}),G=e=>{var a;const s=null==(a=null==e?void 0:e.props)?void 0:a.name;if(!s)return;const l=F.value[s];j.push({name:l.name,query:l.query,params:l.params})},M=e=>{const a=V.value.findIndex((a=>E(a)===e));E(x)===e&&(1===V.value.length?j.push({name:D.value}):a<V.value.length-1?j.push({name:V.value[a+1].name,query:V.value[a+1].query,params:V.value[a+1].params}):j.push({name:V.value[a-1].name,query:V.value[a-1].query,params:V.value[a-1].params})),V.value.splice(a,1)};n((()=>P.value),(()=>{P.value?document.body.addEventListener("click",(()=>{P.value=!1})):document.body.removeEventListener("click",(()=>{P.value=!1}))})),n((()=>x),((e,a)=>{"Login"!==e.name&&"Reload"!==e.name&&(V.value=V.value.filter((e=>!e.meta.closeTab)),B(e),sessionStorage.setItem("historys",JSON.stringify(V.value)),A.value=window.sessionStorage.getItem("activeValue"))}),{deep:!0}),n((()=>V.value),(()=>{sessionStorage.setItem("historys",JSON.stringify(V.value)),F.value={},V.value.forEach((e=>{F.value[E(e)]=e})),u.emit("setKeepAlive",V.value)}),{deep:!0});return(()=>{u.on("closeThisPage",(()=>{M(L(x))})),u.on("closeAllPage",(()=>{K()})),u.on("mobile",(e=>{_.value=e})),u.on("collapse",(e=>{U.value=e}));const e=[{name:D.value,meta:{title:"首页"},query:{},params:{}}];V.value=JSON.parse(sessionStorage.getItem("historys"))||e,window.sessionStorage.getItem("activeValue")?A.value=window.sessionStorage.getItem("activeValue"):A.value=E(x),B(x),"true"===window.sessionStorage.getItem("needCloseAll")&&(K(),window.sessionStorage.removeItem("needCloseAll"))})(),o((()=>{u.off("collapse"),u.off("mobile")})),(e,a)=>{const s=k,l=w;return r(),i("div",C,[v(l,{modelValue:A.value,"onUpdate:modelValue":a[0]||(a[0]=e=>A.value=e),closable:!(1===V.value.length&&e.$route.name===g(D)),type:"card",onContextmenu:a[1]||(a[1]=I((e=>(e=>{if(1===V.value.length&&x.name===D.value)return!1;let a="";if(a="SPAN"===e.srcElement.nodeName?e.srcElement.offsetParent.id:e.srcElement.id,a){let s;P.value=!0,s=U.value?54:220,_.value&&(s=0),R.value=e.clientX-s,H.value=e.clientY+10,z.value=a.substring(4)}})(e)),["prevent"])),onTabClick:G,onTabRemove:M},{default:m((()=>[(r(!0),i(c,null,d(V.value,(e=>(r(),p(s,{key:L(e),label:e.meta.title,name:L(e),tab:e,class:"gva-tab"},{label:m((()=>[y("span",{tab:e,style:f({color:A.value===L(e)?g(T).activeColor:"#333"})},[y("i",{class:"dot",style:f({backgroundColor:A.value===L(e)?g(T).activeColor:"#ddd"})},null,4),h(" "+b(g(S)(e.meta.title,e)),1)],12,N)])),_:2},1032,["label","name","tab"])))),128))])),_:1},8,["modelValue","closable"]),q(y("ul",{style:f({left:R.value+"px",top:H.value+"px"}),class:"contextmenu"},[y("li",{onClick:K},"关闭所有"),y("li",{onClick:X},"关闭左侧"),y("li",{onClick:Y},"关闭右侧"),y("li",{onClick:$},"关闭其他")],4),[[O,P.value]])])}}});export{x as default};