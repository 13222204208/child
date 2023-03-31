/*! 
 Build based on gin-vue-admin 
 Time : 1680141068000 */
import{q as e,A as a,bq as s,aZ as t,av as l,b2 as r,aj as o,z as i,r as d,C as n,ai as c,Z as f,_ as u,aA as p,o as b,a0 as m,e as v,w as y,af as h,d as w,aX as C,ah as k,n as R,c as z,I as E,a3 as S,f as g,ag as A,aD as I,a_ as $,a5 as _}from"./index.9fcd1cf4.js";import{b as L,c as x,u as q}from"./overlay.a3adf688.js";const B=a({name:"ElDrawer",components:{ElOverlay:s,ElFocusTrap:t,ElIcon:l,Close:r},inheritAttrs:!1,props:e({...L,direction:{type:String,default:"rtl",values:["ltr","rtl","ttb","btt"]},size:{type:[String,Number],default:"30%"},withHeader:{type:Boolean,default:!0},modalFade:{type:Boolean,default:!0}}),emits:x,setup(e,{slots:a}){o({scope:"el-drawer",from:"the title slot",replacement:"the header slot",version:"3.0.0",ref:"https://element-plus.org/en-US/component/drawer.html#slots"},i((()=>!!a.title))),o({scope:"el-drawer",from:"custom-class",replacement:"class",version:"2.3.0",ref:"https://element-plus.org/en-US/component/drawer.html#attributes",type:"Attribute"},i((()=>!!e.customClass)));const s=d(),t=d(),l=n("drawer"),{t:r}=c(),u=i((()=>"rtl"===e.direction||"ltr"===e.direction)),p=i((()=>f(e.size)));return{...q(e,s),drawerRef:s,focusStartRef:t,isHorizontal:u,drawerSize:p,ns:l,t:r}}}),D=["aria-label","aria-labelledby","aria-describedby"],H=["id"],T=["aria-label"],j=["id"];const F=_(u(B,[["render",function(e,a,s,t,l,r){const o=p("close"),i=p("el-icon"),d=p("el-focus-trap"),n=p("el-overlay");return b(),m($,{to:"body",disabled:!e.appendToBody},[v(I,{name:e.ns.b("fade"),onAfterEnter:e.afterEnter,onAfterLeave:e.afterLeave,onBeforeLeave:e.beforeLeave,persisted:""},{default:y((()=>[h(v(n,{mask:e.modal,"overlay-class":e.modalClass,"z-index":e.zIndex,onClick:e.onModalClick},{default:y((()=>[v(d,{loop:"",trapped:e.visible,"focus-trap-el":e.drawerRef,"focus-start-el":e.focusStartRef,onReleaseRequested:e.onCloseRequested},{default:y((()=>[w("div",C({ref:"drawerRef","aria-modal":"true","aria-label":e.title||void 0,"aria-labelledby":e.title?void 0:e.titleId,"aria-describedby":e.bodyId},e.$attrs,{class:[e.ns.b(),e.direction,e.visible&&"open",e.customClass],style:e.isHorizontal?"width: "+e.drawerSize:"height: "+e.drawerSize,role:"dialog",onClick:a[1]||(a[1]=k((()=>{}),["stop"]))}),[w("span",{ref:"focusStartRef",class:R(e.ns.e("sr-focus")),tabindex:"-1"},null,2),e.withHeader?(b(),z("header",{key:0,class:R(e.ns.e("header"))},[e.$slots.title?E(e.$slots,"title",{key:1},(()=>[g(" DEPRECATED SLOT ")])):E(e.$slots,"header",{key:0,close:e.handleClose,titleId:e.titleId,titleClass:e.ns.e("title")},(()=>[e.$slots.title?g("v-if",!0):(b(),z("span",{key:0,id:e.titleId,role:"heading",class:R(e.ns.e("title"))},S(e.title),11,H))])),e.showClose?(b(),z("button",{key:2,"aria-label":e.t("el.drawer.close"),class:R(e.ns.e("close-btn")),type:"button",onClick:a[0]||(a[0]=(...a)=>e.handleClose&&e.handleClose(...a))},[v(i,{class:R(e.ns.e("close"))},{default:y((()=>[v(o)])),_:1},8,["class"])],10,T)):g("v-if",!0)],2)):g("v-if",!0),e.rendered?(b(),z("div",{key:1,id:e.bodyId,class:R(e.ns.e("body"))},[E(e.$slots,"default")],10,j)):g("v-if",!0),e.$slots.footer?(b(),z("div",{key:2,class:R(e.ns.e("footer"))},[E(e.$slots,"footer")],2)):g("v-if",!0)],16,D)])),_:3},8,["trapped","focus-trap-el","focus-start-el","onReleaseRequested"])])),_:3},8,["mask","overlay-class","z-index","onClick"]),[[A,e.visible]])])),_:3},8,["name","onAfterEnter","onAfterLeave","onBeforeLeave"])],8,["disabled"])}],["__file","/home/runner/work/element-plus/element-plus/packages/components/drawer/src/drawer.vue"]]));export{F as E};