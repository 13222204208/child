/*! 
 Build based on gin-vue-admin 
 Time : 1679909256000 */
import{q as e,ax as a,s as t,aw as i,aq as l,y as n,x as o,an as c,au as s,cS as r,A as v,al as u,B as d,C as p,aj as f,z as m,bT as y,bS as b,r as h,Z as g,D as k,K as x,P as S,o as V,c as C,d as I,J as w,n as T,aE as B,a0 as _,w as N,a1 as P,av as E,f as K,a3 as j,e as q,bN as z,a2 as A,ah as D,_ as F,ae as J,V as M,cM as O,O as U,a5 as Z}from"./index.61dd0a2d.js";const G=e({modelValue:{type:[Boolean,String,Number],default:!1},value:{type:[Boolean,String,Number],default:!1},disabled:{type:Boolean,default:!1},width:{type:[String,Number],default:""},inlinePrompt:{type:Boolean,default:!1},activeIcon:{type:a},inactiveIcon:{type:a},activeText:{type:String,default:""},inactiveText:{type:String,default:""},activeColor:{type:String,default:""},inactiveColor:{type:String,default:""},borderColor:{type:String,default:""},activeValue:{type:[Boolean,String,Number],default:!0},inactiveValue:{type:[Boolean,String,Number],default:!1},name:{type:String,default:""},validateEvent:{type:Boolean,default:!0},id:String,loading:{type:Boolean,default:!1},beforeChange:{type:t(Function)},size:{type:String,validator:i},tabindex:{type:[String,Number]}}),H={[l]:e=>n(e)||o(e)||c(e),[s]:e=>n(e)||o(e)||c(e),[r]:e=>n(e)||o(e)||c(e)},L=["onClick"],Q=["id","aria-checked","aria-disabled","name","true-value","false-value","disabled","tabindex","onKeydown"],R=["aria-hidden"],W=["aria-hidden"],X=["aria-hidden"],Y="ElSwitch",$=v({name:Y});const ee=Z(F(v({...$,props:G,emits:H,setup(e,{expose:a,emit:t}){const i=e,o=J(),{formItem:c}=u(),v=d(),F=p("switch");f({from:'"value"',replacement:'"model-value" or "v-model"',scope:Y,version:"2.3.0",ref:"https://element-plus.org/en-US/component/switch.html#attributes",type:"Attribute"},m((()=>{var e;return!!(null==(e=o.vnode.props)?void 0:e.value)})));const{inputId:Z}=y(i,{formItemContext:c}),G=b(m((()=>i.loading))),H=h(!1!==i.modelValue),$=h(),ee=h(),ae=m((()=>[F.b(),F.m(v.value),F.is("disabled",G.value),F.is("checked",le.value)])),te=m((()=>({width:g(i.width)})));k((()=>i.modelValue),(()=>{H.value=!0})),k((()=>i.value),(()=>{H.value=!1}));const ie=m((()=>H.value?i.modelValue:i.value)),le=m((()=>ie.value===i.activeValue));[i.activeValue,i.inactiveValue].includes(ie.value)||(t(l,i.inactiveValue),t(s,i.inactiveValue),t(r,i.inactiveValue)),k(le,(e=>{var a;$.value.checked=e,i.validateEvent&&(null==(a=null==c?void 0:c.validate)||a.call(c,"change").catch((e=>x())))}));const ne=()=>{const e=le.value?i.inactiveValue:i.activeValue;t(l,e),t(s,e),t(r,e),M((()=>{$.value.checked=le.value}))},oe=()=>{if(G.value)return;const{beforeChange:e}=i;if(!e)return void ne();const a=e();[O(a),n(a)].includes(!0)||U(Y,"beforeChange must return type `Promise<boolean>` or `boolean`"),O(a)?a.then((e=>{e&&ne()})).catch((e=>{})):a&&ne()},ce=m((()=>F.cssVarBlock({...i.activeColor?{"on-color":i.activeColor}:null,...i.inactiveColor?{"off-color":i.inactiveColor}:null,...i.borderColor?{"border-color":i.borderColor}:null})));return S((()=>{$.value.checked=le.value})),a({focus:()=>{var e,a;null==(a=null==(e=$.value)?void 0:e.focus)||a.call(e)},checked:le}),(e,a)=>(V(),C("div",{class:T(w(ae)),style:A(w(ce)),onClick:D(oe,["prevent"])},[I("input",{id:w(Z),ref_key:"input",ref:$,class:T(w(F).e("input")),type:"checkbox",role:"switch","aria-checked":w(le),"aria-disabled":w(G),name:e.name,"true-value":e.activeValue,"false-value":e.inactiveValue,disabled:w(G),tabindex:e.tabindex,onChange:ne,onKeydown:B(oe,["enter"])},null,42,Q),e.inlinePrompt||!e.inactiveIcon&&!e.inactiveText?K("v-if",!0):(V(),C("span",{key:0,class:T([w(F).e("label"),w(F).em("label","left"),w(F).is("active",!w(le))])},[e.inactiveIcon?(V(),_(w(E),{key:0},{default:N((()=>[(V(),_(P(e.inactiveIcon)))])),_:1})):K("v-if",!0),!e.inactiveIcon&&e.inactiveText?(V(),C("span",{key:1,"aria-hidden":w(le)},j(e.inactiveText),9,R)):K("v-if",!0)],2)),I("span",{ref_key:"core",ref:ee,class:T(w(F).e("core")),style:A(w(te))},[e.inlinePrompt?(V(),C("div",{key:0,class:T(w(F).e("inner"))},[e.activeIcon||e.inactiveIcon?(V(),_(w(E),{key:0,class:T(w(F).is("icon"))},{default:N((()=>[(V(),_(P(w(le)?e.activeIcon:e.inactiveIcon)))])),_:1},8,["class"])):e.activeText||e.inactiveText?(V(),C("span",{key:1,class:T(w(F).is("text")),"aria-hidden":!w(le)},j(w(le)?e.activeText:e.inactiveText),11,W)):K("v-if",!0)],2)):K("v-if",!0),I("div",{class:T(w(F).e("action"))},[e.loading?(V(),_(w(E),{key:0,class:T(w(F).is("loading"))},{default:N((()=>[q(w(z))])),_:1},8,["class"])):K("v-if",!0)],2)],6),e.inlinePrompt||!e.activeIcon&&!e.activeText?K("v-if",!0):(V(),C("span",{key:1,class:T([w(F).e("label"),w(F).em("label","right"),w(F).is("active",w(le))])},[e.activeIcon?(V(),_(w(E),{key:0},{default:N((()=>[(V(),_(P(e.activeIcon)))])),_:1})):K("v-if",!0),!e.activeIcon&&e.activeText?(V(),C("span",{key:1,"aria-hidden":!w(le)},j(e.activeText),9,X)):K("v-if",!0)],2))],14,L))}}),[["__file","/home/runner/work/element-plus/element-plus/packages/components/switch/src/switch.vue"]]));export{ee as E};
