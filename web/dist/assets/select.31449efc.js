/*! 
 Build based on gin-vue-admin 
 Time : 1679909256000 */
import{ab as e,M as t,z as l,ac as o,ad as a,D as n,J as s,ae as i,A as r,C as u,a as p,G as d,Q as c,V as v,_ as h,af as m,ag as f,o as b,c as g,I as y,d as C,a3 as S,n as x,ah as O,r as w,P as L,T as I,a2 as M,ai as T,aj as E,ak as V,al as P,B as k,K as B,L as z,am as D,a9 as q,an as W,ao as F,ap as A,aq as K,x as H,ar as _,as as R,at as $,au as j,b as Q,av as N,aw as G,ax as U,ay as J,az as X,F as Y,aA as Z,aB as ee,e as te,w as le,a0 as oe,U as ae,aC as ne,f as se,aD as ie,aE as re,aF as ue,aG as pe,a1 as de,a5 as ce,a6 as ve}from"./index.61dd0a2d.js";import{E as he,u as me,a as fe}from"./index.cbcb4c8f.js";import{E as be}from"./scrollbar.ef9dc6d6.js";import{E as ge,t as ye}from"./index.ba9b69da.js";import{e as Ce}from"./strings.ed5ba909.js";import{i as Se}from"./isEqual.3dfeee36.js";import{d as xe,C as Oe}from"./tag.3ecd1fed.js";const we=Symbol("ElSelectGroup"),Le=Symbol("ElSelect");var Ie=h(r({name:"ElOption",componentName:"ElOption",props:{value:{required:!0,type:[String,Number,Boolean,Object]},label:[String,Number],created:Boolean,disabled:{type:Boolean,default:!1}},setup(e){const r=u("select"),h=p({index:-1,groupDisabled:!1,visible:!0,hitState:!1,hover:!1}),{currentLabel:m,itemSelected:f,isDisabled:b,select:g,hoverItem:y}=function(e,r){const u=t(Le),p=t(we,{disabled:!1}),d=l((()=>"[object object]"===Object.prototype.toString.call(e.value).toLowerCase())),c=l((()=>u.props.multiple?g(u.props.modelValue,e.value):y(e.value,u.props.modelValue))),v=l((()=>{if(u.props.multiple){const e=u.props.modelValue||[];return!c.value&&e.length>=u.props.multipleLimit&&u.props.multipleLimit>0}return!1})),h=l((()=>e.label||(d.value?"":e.value))),m=l((()=>e.value||e.label||"")),f=l((()=>e.disabled||r.groupDisabled||v.value)),b=i(),g=(e=[],t)=>{if(d.value){const l=u.props.valueKey;return e&&e.some((e=>o(a(e,l))===a(t,l)))}return e&&e.includes(t)},y=(e,t)=>{if(d.value){const{valueKey:l}=u.props;return a(e,l)===a(t,l)}return e===t};n((()=>h.value),(()=>{e.created||u.props.remote||u.setSelected()})),n((()=>e.value),((t,l)=>{const{remote:o,valueKey:a}=u.props;if(Object.is(t,l)||(u.onOptionDestroy(l,b.proxy),u.onOptionCreate(b.proxy)),!e.created&&!o){if(a&&"object"==typeof t&&"object"==typeof l&&t[a]===l[a])return;u.setSelected()}})),n((()=>p.disabled),(()=>{r.groupDisabled=p.disabled}),{immediate:!0});const{queryChange:C}=o(u);return n(C,(t=>{const{query:l}=s(t),o=new RegExp(Ce(l),"i");r.visible=o.test(h.value)||e.created,r.visible||u.filteredOptionsCount--})),{select:u,currentLabel:h,currentValue:m,itemSelected:c,isDisabled:f,hoverItem:()=>{e.disabled||p.disabled||(u.hoverIndex=u.optionsArray.indexOf(b.proxy))}}}(e,h),{visible:C,hover:S}=d(h),x=i().proxy;return g.onOptionCreate(x),c((()=>{const e=x.value,{selected:t}=g,l=(g.props.multiple?t:[t]).some((e=>e.value===x.value));v((()=>{g.cachedOptions.get(e)!==x||l||g.cachedOptions.delete(e)})),g.onOptionDestroy(e,x)})),{ns:r,currentLabel:m,itemSelected:f,isDisabled:b,select:g,hoverItem:y,visible:C,hover:S,selectOptionClick:function(){!0!==e.disabled&&!0!==h.groupDisabled&&g.handleOptionSelect(x,!0)},states:h}}}),[["render",function(e,t,l,o,a,n){return m((b(),g("li",{class:x([e.ns.be("dropdown","item"),e.ns.is("disabled",e.isDisabled),{selected:e.itemSelected,hover:e.hover}]),onMouseenter:t[0]||(t[0]=(...t)=>e.hoverItem&&e.hoverItem(...t)),onClick:t[1]||(t[1]=O(((...t)=>e.selectOptionClick&&e.selectOptionClick(...t)),["stop"]))},[y(e.$slots,"default",{},(()=>[C("span",null,S(e.currentLabel),1)]))],34)),[[f,e.visible]])}],["__file","/home/runner/work/element-plus/element-plus/packages/components/select/src/option.vue"]]);const Me=r({name:"ElSelectDropdown",componentName:"ElSelectDropdown",setup(){const e=t(Le),o=u("select"),a=l((()=>e.props.popperClass)),n=l((()=>e.props.multiple)),s=l((()=>e.props.fitInputWidth)),i=w("");function r(){var t;i.value=`${null==(t=e.selectWrapper)?void 0:t.offsetWidth}px`}return L((()=>{r(),I(e.selectWrapper,r)})),{ns:o,minWidth:i,popperClass:a,isMultiple:n,isFitInputWidth:s}}});const Te=(t,s,i)=>{const{t:r}=T(),p=u("select");E({from:"suffixTransition",replacement:"override style scheme",version:"2.3.0",scope:"props",ref:"https://element-plus.org/en-US/component/select.html#select-attributes"},l((()=>!1===t.suffixTransition)));const d=w(null),c=w(null),h=w(null),m=w(null),f=w(null),b=w(null),g=w(-1),y=V({query:""}),C=V(""),{form:S,formItem:x}=P(),O=l((()=>!t.filterable||t.multiple||!s.visible)),L=l((()=>t.disabled||(null==S?void 0:S.disabled))),I=l((()=>{const e=t.multiple?Array.isArray(t.modelValue)&&t.modelValue.length>0:void 0!==t.modelValue&&null!==t.modelValue&&""!==t.modelValue;return t.clearable&&!L.value&&s.inputHovering&&e})),M=l((()=>t.remote&&t.filterable&&!t.remoteShowSuffix?"":t.suffixIcon)),Q=l((()=>p.is("reverse",M.value&&s.visible&&t.suffixTransition))),N=l((()=>t.remote?300:0)),G=l((()=>t.loading?t.loadingText||r("el.select.loading"):(!t.remote||""!==s.query||0!==s.options.size)&&(t.filterable&&s.query&&s.options.size>0&&0===s.filteredOptionsCount?t.noMatchText||r("el.select.noMatch"):0===s.options.size?t.noDataText||r("el.select.noData"):null))),U=l((()=>Array.from(s.options.values()))),J=l((()=>Array.from(s.cachedOptions.values()))),X=l((()=>{const e=U.value.filter((e=>!e.created)).some((e=>e.currentLabel===s.query));return t.filterable&&t.allowCreate&&""!==s.query&&!e})),Y=k(),Z=l((()=>["small"].includes(Y.value)?"small":"default")),ee=l({get:()=>s.visible&&!1!==G.value,set(e){s.visible=e}});n([()=>L.value,()=>Y.value,()=>null==S?void 0:S.size],(()=>{v((()=>{te()}))})),n((()=>t.placeholder),(e=>{s.cachedPlaceHolder=s.currentPlaceholder=e})),n((()=>t.modelValue),((e,l)=>{t.multiple&&(te(),e&&e.length>0||c.value&&""!==s.query?s.currentPlaceholder="":s.currentPlaceholder=s.cachedPlaceHolder,t.filterable&&!t.reserveKeyword&&(s.query="",le(s.query))),ne(),t.filterable&&!t.multiple&&(s.inputLength=20),!Se(e,l)&&t.validateEvent&&(null==x||x.validate("change").catch((e=>B())))}),{flush:"post",deep:!0}),n((()=>s.visible),(e=>{var l,o,a;e?(null==(o=null==(l=h.value)?void 0:l.updatePopper)||o.call(l),t.filterable&&(s.filteredOptionsCount=s.optionsCount,s.query=t.remote?"":s.selectedLabel,t.multiple?null==(a=c.value)||a.focus():s.selectedLabel&&(s.currentPlaceholder=`${s.selectedLabel}`,s.selectedLabel=""),le(s.query),t.multiple||t.remote||(y.value.query="",D(y),D(C)))):(t.filterable&&(z(t.filterMethod)&&t.filterMethod(""),z(t.remoteMethod)&&t.remoteMethod("")),c.value&&c.value.blur(),s.query="",s.previousQuery=null,s.selectedLabel="",s.inputLength=20,s.menuVisibleOnFocus=!1,ie(),v((()=>{c.value&&""===c.value.value&&0===s.selected.length&&(s.currentPlaceholder=s.cachedPlaceHolder)})),t.multiple||(s.selected&&(t.filterable&&t.allowCreate&&s.createdSelected&&s.createdLabel?s.selectedLabel=s.createdLabel:s.selectedLabel=s.selected.currentLabel,t.filterable&&(s.query=s.selectedLabel)),t.filterable&&(s.currentPlaceholder=s.cachedPlaceHolder))),i.emit("visible-change",e)})),n((()=>s.options.entries()),(()=>{var e,l,o;if(!q)return;null==(l=null==(e=h.value)?void 0:e.updatePopper)||l.call(e),t.multiple&&te();const a=(null==(o=f.value)?void 0:o.querySelectorAll("input"))||[];Array.from(a).includes(document.activeElement)||ne(),t.defaultFirstOption&&(t.filterable||t.remote)&&s.filteredOptionsCount&&ae()}),{flush:"post"}),n((()=>s.hoverIndex),(e=>{W(e)&&e>-1?g.value=U.value[e]||{}:g.value={},U.value.forEach((e=>{e.hover=g.value===e}))}));const te=()=>{t.collapseTags&&!t.filterable||v((()=>{var t,l;if(!d.value)return;const o=d.value.$el.querySelector("input"),a=m.value,n=(i=Y.value||(null==S?void 0:S.size),e[i||"default"]);var i;o.style.height=(0===s.selected.length?n:Math.max(a?a.clientHeight+(a.clientHeight>n?6:0):0,n))-2+"px",s.tagInMultiLine=Number.parseFloat(o.style.height)>=n,s.visible&&!1!==G.value&&(null==(l=null==(t=h.value)?void 0:t.updatePopper)||l.call(t))}))},le=async e=>{s.previousQuery===e||s.isOnComposition||(null!==s.previousQuery||!z(t.filterMethod)&&!z(t.remoteMethod)?(s.previousQuery=e,v((()=>{var e,t;s.visible&&(null==(t=null==(e=h.value)?void 0:e.updatePopper)||t.call(e))})),s.hoverIndex=-1,t.multiple&&t.filterable&&v((()=>{const e=15*c.value.value.length+20;s.inputLength=t.collapseTags?Math.min(50,e):e,oe(),te()})),t.remote&&z(t.remoteMethod)?(s.hoverIndex=-1,t.remoteMethod(e)):z(t.filterMethod)?(t.filterMethod(e),D(C)):(s.filteredOptionsCount=s.optionsCount,y.value.query=e,D(y),D(C)),t.defaultFirstOption&&(t.filterable||t.remote)&&s.filteredOptionsCount&&(await v(),ae())):s.previousQuery=e)},oe=()=>{""!==s.currentPlaceholder&&(s.currentPlaceholder=c.value.value?"":s.cachedPlaceHolder)},ae=()=>{const e=U.value.filter((e=>e.visible&&!e.disabled&&!e.states.groupDisabled)),t=e.find((e=>e.created)),l=e[0];s.hoverIndex=he(U.value,t||l)},ne=()=>{var e;if(!t.multiple){const l=se(t.modelValue);return(null==(e=l.props)?void 0:e.created)?(s.createdLabel=l.props.value,s.createdSelected=!0):s.createdSelected=!1,s.selectedLabel=l.currentLabel,s.selected=l,void(t.filterable&&(s.query=s.selectedLabel))}s.selectedLabel="";const l=[];Array.isArray(t.modelValue)&&t.modelValue.forEach((e=>{l.push(se(e))})),s.selected=l,v((()=>{te()}))},se=e=>{let l;const o="object"===F(e).toLowerCase(),n="null"===F(e).toLowerCase(),i="undefined"===F(e).toLowerCase();for(let u=s.cachedOptions.size-1;u>=0;u--){const n=J.value[u];if(o?a(n.value,t.valueKey)===a(e,t.valueKey):n.value===e){l={value:e,currentLabel:n.currentLabel,isDisabled:n.isDisabled};break}}if(l)return l;const r={value:e,currentLabel:o?e.label:n||i?"":e};return t.multiple&&(r.hitState=!1),r},ie=()=>{setTimeout((()=>{const e=t.valueKey;t.multiple?s.selected.length>0?s.hoverIndex=Math.min.apply(null,s.selected.map((t=>U.value.findIndex((l=>a(l,e)===a(t,e)))))):s.hoverIndex=-1:s.hoverIndex=U.value.findIndex((e=>ye(e)===ye(s.selected)))}),300)},re=()=>{var e;s.inputWidth=null==(e=d.value)?void 0:e.$el.offsetWidth},ue=xe((()=>{t.filterable&&s.query!==s.selectedLabel&&(s.query=s.selectedLabel,le(s.query))}),N.value),pe=xe((e=>{le(e.target.value)}),N.value),de=e=>{Se(t.modelValue,e)||i.emit(j,e)},ce=e=>{e.stopPropagation();const l=t.multiple?[]:"";if(!H(l))for(const t of s.selected)t.isDisabled&&l.push(t.value);i.emit(K,l),de(l),s.hoverIndex=-1,s.visible=!1,i.emit("clear")},ve=(e,l)=>{var o;if(t.multiple){const l=(t.modelValue||[]).slice(),a=he(l,e.value);a>-1?l.splice(a,1):(t.multipleLimit<=0||l.length<t.multipleLimit)&&l.push(e.value),i.emit(K,l),de(l),e.created&&(s.query="",le(""),s.inputLength=20),t.filterable&&(null==(o=c.value)||o.focus())}else i.emit(K,e.value),de(e.value),s.visible=!1;s.isSilentBlur=l,me(),s.visible||v((()=>{fe(e)}))},he=(e=[],l)=>{if(!A(l))return e.indexOf(l);const n=t.valueKey;let s=-1;return e.some(((e,t)=>o(a(e,n))===a(l,n)&&(s=t,!0))),s},me=()=>{s.softFocus=!0;const e=c.value||d.value;e&&(null==e||e.focus())},fe=e=>{var t,l,o,a,n;const s=Array.isArray(e)?e[0]:e;let i=null;if(null==s?void 0:s.value){const e=U.value.filter((e=>e.value===s.value));e.length>0&&(i=e[0].$el)}if(h.value&&i){const e=null==(a=null==(o=null==(l=null==(t=h.value)?void 0:t.popperRef)?void 0:l.contentRef)?void 0:o.querySelector)?void 0:a.call(o,`.${p.be("dropdown","wrap")}`);e&&_(e,i)}null==(n=b.value)||n.handleScroll()},be=e=>{if(!Array.isArray(s.selected))return;const t=s.selected[s.selected.length-1];return t?!0===e||!1===e?(t.hitState=e,e):(t.hitState=!t.hitState,t.hitState):void 0},ge=e=>{var t;e&&!s.mouseEnter||L.value||(s.menuVisibleOnFocus?s.menuVisibleOnFocus=!1:h.value&&h.value.isFocusInsideContent()||(s.visible=!s.visible),s.visible&&(null==(t=c.value||d.value)||t.focus()))},ye=e=>A(e.value)?a(e.value,t.valueKey):e.value,Ce=l((()=>U.value.filter((e=>e.visible)).every((e=>e.disabled)))),Oe=e=>{if(s.visible){if(0!==s.options.size&&0!==s.filteredOptionsCount&&!s.isOnComposition&&!Ce.value){"next"===e?(s.hoverIndex++,s.hoverIndex===s.options.size&&(s.hoverIndex=0)):"prev"===e&&(s.hoverIndex--,s.hoverIndex<0&&(s.hoverIndex=s.options.size-1));const t=U.value[s.hoverIndex];!0!==t.disabled&&!0!==t.states.groupDisabled&&t.visible||Oe(e),v((()=>fe(g.value)))}}else s.visible=!0};return{optionsArray:U,selectSize:Y,handleResize:()=>{var e,l;re(),null==(l=null==(e=h.value)?void 0:e.updatePopper)||l.call(e),t.multiple&&te()},debouncedOnInputChange:ue,debouncedQueryChange:pe,deletePrevTag:e=>{if(e.target.value.length<=0&&!be()){const e=t.modelValue.slice();e.pop(),i.emit(K,e),de(e)}1===e.target.value.length&&0===t.modelValue.length&&(s.currentPlaceholder=s.cachedPlaceHolder)},deleteTag:(e,l)=>{const o=s.selected.indexOf(l);if(o>-1&&!L.value){const e=t.modelValue.slice();e.splice(o,1),i.emit(K,e),de(e),i.emit("remove-tag",l.value)}e.stopPropagation()},deleteSelected:ce,handleOptionSelect:ve,scrollToOption:fe,readonly:O,resetInputHeight:te,showClose:I,iconComponent:M,iconReverse:Q,showNewOption:X,collapseTagSize:Z,setSelected:ne,managePlaceholder:oe,selectDisabled:L,emptyText:G,toggleLastOptionHitState:be,resetInputState:e=>{e.code!==R.backspace&&be(!1),s.inputLength=15*c.value.value.length+20,te()},handleComposition:e=>{const t=e.target.value;if("compositionend"===e.type)s.isOnComposition=!1,v((()=>le(t)));else{const e=t[t.length-1]||"";s.isOnComposition=!$(e)}},onOptionCreate:e=>{s.optionsCount++,s.filteredOptionsCount++,s.options.set(e.value,e),s.cachedOptions.set(e.value,e)},onOptionDestroy:(e,t)=>{s.options.get(e)===t&&(s.optionsCount--,s.filteredOptionsCount--,s.options.delete(e))},handleMenuEnter:()=>{v((()=>fe(s.selected)))},handleFocus:e=>{s.softFocus?s.softFocus=!1:((t.automaticDropdown||t.filterable)&&(t.filterable&&!s.visible&&(s.menuVisibleOnFocus=!0),s.visible=!0),i.emit("focus",e))},blur:()=>{var e;s.visible=!1,null==(e=d.value)||e.blur()},handleBlur:e=>{v((()=>{s.isSilentBlur?s.isSilentBlur=!1:i.emit("blur",e)})),s.softFocus=!1},handleClearClick:e=>{ce(e)},handleClose:()=>{s.visible=!1},handleKeydownEscape:e=>{s.visible&&(e.preventDefault(),e.stopPropagation(),s.visible=!1)},toggleMenu:ge,selectOption:()=>{s.visible?U.value[s.hoverIndex]&&ve(U.value[s.hoverIndex],void 0):ge()},getValueKey:ye,navigateOptions:Oe,dropMenuVisible:ee,queryChange:y,groupQueryChange:C,reference:d,input:c,tooltipRef:h,tags:m,selectWrapper:f,scrollbar:b,handleMouseEnter:()=>{s.mouseEnter=!0},handleMouseLeave:()=>{s.mouseEnter=!1}}},Ee="ElSelect",Ve=r({name:Ee,componentName:Ee,components:{ElInput:Q,ElSelectMenu:h(Me,[["render",function(e,t,l,o,a,n){return b(),g("div",{class:x([e.ns.b("dropdown"),e.ns.is("multiple",e.isMultiple),e.popperClass]),style:M({[e.isFitInputWidth?"width":"minWidth"]:e.minWidth})},[y(e.$slots,"default")],6)}],["__file","/home/runner/work/element-plus/element-plus/packages/components/select/src/select-dropdown.vue"]]),ElOption:Ie,ElTag:ge,ElScrollbar:be,ElTooltip:he,ElIcon:N},directives:{ClickOutside:Oe},props:{name:String,id:String,modelValue:{type:[Array,String,Number,Boolean,Object],default:void 0},autocomplete:{type:String,default:"off"},automaticDropdown:Boolean,size:{type:String,validator:G},effect:{type:String,default:"light"},disabled:Boolean,clearable:Boolean,filterable:Boolean,allowCreate:Boolean,loading:Boolean,popperClass:{type:String,default:""},remote:Boolean,loadingText:String,noMatchText:String,noDataText:String,remoteMethod:Function,filterMethod:Function,multiple:Boolean,multipleLimit:{type:Number,default:0},placeholder:{type:String},defaultFirstOption:Boolean,reserveKeyword:{type:Boolean,default:!0},valueKey:{type:String,default:"value"},collapseTags:Boolean,collapseTagsTooltip:{type:Boolean,default:!1},teleported:me.teleported,persistent:{type:Boolean,default:!0},clearIcon:{type:U,default:J},fitInputWidth:{type:Boolean,default:!1},suffixIcon:{type:U,default:X},tagType:{...ye.type,default:"info"},validateEvent:{type:Boolean,default:!0},remoteShowSuffix:{type:Boolean,default:!1},suffixTransition:{type:Boolean,default:!0},placement:{type:String,values:fe,default:"bottom-start"}},emits:[K,j,"remove-tag","clear","visible-change","focus","blur"],setup(e,t){const o=u("select"),a=u("input"),{t:n}=T(),i=function(e){const{t:t}=T();return p({options:new Map,cachedOptions:new Map,createdLabel:null,createdSelected:!1,selected:e.multiple?[]:{},inputLength:20,inputWidth:0,optionsCount:0,filteredOptionsCount:0,visible:!1,softFocus:!1,selectedLabel:"",hoverIndex:-1,query:"",previousQuery:null,inputHovering:!1,cachedPlaceHolder:"",currentPlaceholder:t("el.select.placeholder"),menuVisibleOnFocus:!1,isOnComposition:!1,isSilentBlur:!1,prefixWidth:11,tagInMultiLine:!1,mouseEnter:!1})}(e),{optionsArray:r,selectSize:c,readonly:h,handleResize:m,collapseTagSize:f,debouncedOnInputChange:b,debouncedQueryChange:g,deletePrevTag:y,deleteTag:C,deleteSelected:S,handleOptionSelect:x,scrollToOption:O,setSelected:w,resetInputHeight:M,managePlaceholder:E,showClose:V,selectDisabled:P,iconComponent:k,iconReverse:B,showNewOption:z,emptyText:D,toggleLastOptionHitState:q,resetInputState:W,handleComposition:F,onOptionCreate:A,onOptionDestroy:H,handleMenuEnter:_,handleFocus:R,blur:$,handleBlur:j,handleClearClick:Q,handleClose:N,handleKeydownEscape:G,toggleMenu:U,selectOption:J,getValueKey:X,navigateOptions:Z,dropMenuVisible:ee,reference:te,input:le,tooltipRef:oe,tags:ae,selectWrapper:ne,scrollbar:se,queryChange:ie,groupQueryChange:re,handleMouseEnter:ue,handleMouseLeave:pe}=Te(e,i,t),{focus:de}=(ce=te,{focus:()=>{var e,t;null==(t=null==(e=ce.value)?void 0:e.focus)||t.call(e)}});var ce;const{inputWidth:ve,selected:he,inputLength:me,filteredOptionsCount:fe,visible:be,softFocus:ge,selectedLabel:ye,hoverIndex:Ce,query:Se,inputHovering:xe,currentPlaceholder:Oe,menuVisibleOnFocus:we,isOnComposition:Ie,isSilentBlur:Me,options:Ee,cachedOptions:Ve,optionsCount:Pe,prefixWidth:ke,tagInMultiLine:Be}=d(i),ze=l((()=>{const t=[o.b()],l=s(c);return l&&t.push(o.m(l)),e.disabled&&t.push(o.m("disabled")),t})),De=l((()=>({maxWidth:s(ve)-32+"px",width:"100%"}))),qe=l((()=>({maxWidth:`${s(ve)>123?s(ve)-123:s(ve)-75}px`})));Y(Le,p({props:e,options:Ee,optionsArray:r,cachedOptions:Ve,optionsCount:Pe,filteredOptionsCount:fe,hoverIndex:Ce,handleOptionSelect:x,onOptionCreate:A,onOptionDestroy:H,selectWrapper:ne,selected:he,setSelected:w,queryChange:ie,groupQueryChange:re})),L((()=>{i.cachedPlaceHolder=Oe.value=e.placeholder||n("el.select.placeholder"),e.multiple&&Array.isArray(e.modelValue)&&e.modelValue.length>0&&(Oe.value=""),I(ne,m),e.remote&&e.multiple&&M(),v((()=>{const e=te.value&&te.value.$el;if(e&&(ve.value=e.getBoundingClientRect().width,t.slots.prefix)){const t=e.querySelector(`.${a.e("prefix")}`);ke.value=Math.max(t.getBoundingClientRect().width+5,30)}})),w()})),e.multiple&&!Array.isArray(e.modelValue)&&t.emit(K,[]),!e.multiple&&Array.isArray(e.modelValue)&&t.emit(K,"");const We=l((()=>{var e,t;return null==(t=null==(e=oe.value)?void 0:e.popperRef)?void 0:t.contentRef}));return{tagInMultiLine:Be,prefixWidth:ke,selectSize:c,readonly:h,handleResize:m,collapseTagSize:f,debouncedOnInputChange:b,debouncedQueryChange:g,deletePrevTag:y,deleteTag:C,deleteSelected:S,handleOptionSelect:x,scrollToOption:O,inputWidth:ve,selected:he,inputLength:me,filteredOptionsCount:fe,visible:be,softFocus:ge,selectedLabel:ye,hoverIndex:Ce,query:Se,inputHovering:xe,currentPlaceholder:Oe,menuVisibleOnFocus:we,isOnComposition:Ie,isSilentBlur:Me,options:Ee,resetInputHeight:M,managePlaceholder:E,showClose:V,selectDisabled:P,iconComponent:k,iconReverse:B,showNewOption:z,emptyText:D,toggleLastOptionHitState:q,resetInputState:W,handleComposition:F,handleMenuEnter:_,handleFocus:R,blur:$,handleBlur:j,handleClearClick:Q,handleClose:N,handleKeydownEscape:G,toggleMenu:U,selectOption:J,getValueKey:X,navigateOptions:Z,dropMenuVisible:ee,focus:de,reference:te,input:le,tooltipRef:oe,popperPaneRef:We,tags:ae,selectWrapper:ne,scrollbar:se,wrapperKls:ze,selectTagsStyle:De,nsSelect:o,tagTextStyle:qe,handleMouseEnter:ue,handleMouseLeave:pe}}}),Pe=["disabled","autocomplete"],ke={style:{height:"100%",display:"flex","justify-content":"center","align-items":"center"}};var Be=h(Ve,[["render",function(e,t,l,o,a,n){const s=Z("el-tag"),i=Z("el-tooltip"),r=Z("el-icon"),u=Z("el-input"),p=Z("el-option"),d=Z("el-scrollbar"),c=Z("el-select-menu"),v=ee("click-outside");return m((b(),g("div",{ref:"selectWrapper",class:x(e.wrapperKls),onMouseenter:t[22]||(t[22]=(...t)=>e.handleMouseEnter&&e.handleMouseEnter(...t)),onMouseleave:t[23]||(t[23]=(...t)=>e.handleMouseLeave&&e.handleMouseLeave(...t)),onClick:t[24]||(t[24]=O(((...t)=>e.toggleMenu&&e.toggleMenu(...t)),["stop"]))},[te(i,{ref:"tooltipRef",visible:e.dropMenuVisible,placement:e.placement,teleported:e.teleported,"popper-class":[e.nsSelect.e("popper"),e.popperClass],"fallback-placements":["bottom-start","top-start","right","left"],effect:e.effect,pure:"",trigger:"click",transition:`${e.nsSelect.namespace.value}-zoom-in-top`,"stop-popper-mouse-event":!1,"gpu-acceleration":!1,persistent:e.persistent,onShow:e.handleMenuEnter},{default:le((()=>[C("div",{class:"select-trigger",onMouseenter:t[20]||(t[20]=t=>e.inputHovering=!0),onMouseleave:t[21]||(t[21]=t=>e.inputHovering=!1)},[e.multiple?(b(),g("div",{key:0,ref:"tags",class:x(e.nsSelect.e("tags")),style:M(e.selectTagsStyle)},[e.collapseTags&&e.selected.length?(b(),g("span",{key:0,class:x([e.nsSelect.b("tags-wrapper"),{"has-prefix":e.prefixWidth&&e.selected.length}])},[te(s,{closable:!e.selectDisabled&&!e.selected[0].isDisabled,size:e.collapseTagSize,hit:e.selected[0].hitState,type:e.tagType,"disable-transitions":"",onClose:t[0]||(t[0]=t=>e.deleteTag(t,e.selected[0]))},{default:le((()=>[C("span",{class:x(e.nsSelect.e("tags-text")),style:M(e.tagTextStyle)},S(e.selected[0].currentLabel),7)])),_:1},8,["closable","size","hit","type"]),e.selected.length>1?(b(),oe(s,{key:0,closable:!1,size:e.collapseTagSize,type:e.tagType,"disable-transitions":""},{default:le((()=>[e.collapseTagsTooltip?(b(),oe(i,{key:0,disabled:e.dropMenuVisible,"fallback-placements":["bottom","top","right","left"],effect:e.effect,placement:"bottom",teleported:e.teleported},{default:le((()=>[C("span",{class:x(e.nsSelect.e("tags-text"))},"+ "+S(e.selected.length-1),3)])),content:le((()=>[C("div",{class:x(e.nsSelect.e("collapse-tags"))},[(b(!0),g(ae,null,ne(e.selected.slice(1),((t,l)=>(b(),g("div",{key:l,class:x(e.nsSelect.e("collapse-tag"))},[(b(),oe(s,{key:e.getValueKey(t),class:"in-tooltip",closable:!e.selectDisabled&&!t.isDisabled,size:e.collapseTagSize,hit:t.hitState,type:e.tagType,"disable-transitions":"",style:{margin:"2px"},onClose:l=>e.deleteTag(l,t)},{default:le((()=>[C("span",{class:x(e.nsSelect.e("tags-text")),style:M({maxWidth:e.inputWidth-75+"px"})},S(t.currentLabel),7)])),_:2},1032,["closable","size","hit","type","onClose"]))],2)))),128))],2)])),_:1},8,["disabled","effect","teleported"])):(b(),g("span",{key:1,class:x(e.nsSelect.e("tags-text"))},"+ "+S(e.selected.length-1),3))])),_:1},8,["size","type"])):se("v-if",!0)],2)):se("v-if",!0),se(" <div> "),e.collapseTags?se("v-if",!0):(b(),oe(ie,{key:1,onAfterLeave:e.resetInputHeight},{default:le((()=>[C("span",{class:x([e.nsSelect.b("tags-wrapper"),{"has-prefix":e.prefixWidth&&e.selected.length}])},[(b(!0),g(ae,null,ne(e.selected,(t=>(b(),oe(s,{key:e.getValueKey(t),closable:!e.selectDisabled&&!t.isDisabled,size:e.collapseTagSize,hit:t.hitState,type:e.tagType,"disable-transitions":"",onClose:l=>e.deleteTag(l,t)},{default:le((()=>[C("span",{class:x(e.nsSelect.e("tags-text")),style:M({maxWidth:e.inputWidth-75+"px"})},S(t.currentLabel),7)])),_:2},1032,["closable","size","hit","type","onClose"])))),128))],2)])),_:1},8,["onAfterLeave"])),se(" </div> "),e.filterable?m((b(),g("input",{key:2,ref:"input","onUpdate:modelValue":t[1]||(t[1]=t=>e.query=t),type:"text",class:x([e.nsSelect.e("input"),e.nsSelect.is(e.selectSize)]),disabled:e.selectDisabled,autocomplete:e.autocomplete,style:M({marginLeft:e.prefixWidth&&!e.selected.length||e.tagInMultiLine?`${e.prefixWidth}px`:"",flexGrow:1,width:e.inputLength/(e.inputWidth-32)+"%",maxWidth:e.inputWidth-42+"px"}),onFocus:t[2]||(t[2]=(...t)=>e.handleFocus&&e.handleFocus(...t)),onBlur:t[3]||(t[3]=(...t)=>e.handleBlur&&e.handleBlur(...t)),onKeyup:t[4]||(t[4]=(...t)=>e.managePlaceholder&&e.managePlaceholder(...t)),onKeydown:[t[5]||(t[5]=(...t)=>e.resetInputState&&e.resetInputState(...t)),t[6]||(t[6]=re(O((t=>e.navigateOptions("next")),["prevent"]),["down"])),t[7]||(t[7]=re(O((t=>e.navigateOptions("prev")),["prevent"]),["up"])),t[8]||(t[8]=re(((...t)=>e.handleKeydownEscape&&e.handleKeydownEscape(...t)),["esc"])),t[9]||(t[9]=re(O(((...t)=>e.selectOption&&e.selectOption(...t)),["stop","prevent"]),["enter"])),t[10]||(t[10]=re(((...t)=>e.deletePrevTag&&e.deletePrevTag(...t)),["delete"])),t[11]||(t[11]=re((t=>e.visible=!1),["tab"]))],onCompositionstart:t[12]||(t[12]=(...t)=>e.handleComposition&&e.handleComposition(...t)),onCompositionupdate:t[13]||(t[13]=(...t)=>e.handleComposition&&e.handleComposition(...t)),onCompositionend:t[14]||(t[14]=(...t)=>e.handleComposition&&e.handleComposition(...t)),onInput:t[15]||(t[15]=(...t)=>e.debouncedQueryChange&&e.debouncedQueryChange(...t))},null,46,Pe)),[[ue,e.query]]):se("v-if",!0)],6)):se("v-if",!0),te(u,{id:e.id,ref:"reference",modelValue:e.selectedLabel,"onUpdate:modelValue":t[16]||(t[16]=t=>e.selectedLabel=t),type:"text",placeholder:e.currentPlaceholder,name:e.name,autocomplete:e.autocomplete,size:e.selectSize,disabled:e.selectDisabled,readonly:e.readonly,"validate-event":!1,class:x([e.nsSelect.is("focus",e.visible)]),tabindex:e.multiple&&e.filterable?-1:void 0,onFocus:e.handleFocus,onBlur:e.handleBlur,onInput:e.debouncedOnInputChange,onPaste:e.debouncedOnInputChange,onCompositionstart:e.handleComposition,onCompositionupdate:e.handleComposition,onCompositionend:e.handleComposition,onKeydown:[t[17]||(t[17]=re(O((t=>e.navigateOptions("next")),["stop","prevent"]),["down"])),t[18]||(t[18]=re(O((t=>e.navigateOptions("prev")),["stop","prevent"]),["up"])),re(O(e.selectOption,["stop","prevent"]),["enter"]),re(e.handleKeydownEscape,["esc"]),t[19]||(t[19]=re((t=>e.visible=!1),["tab"]))]},pe({suffix:le((()=>[e.iconComponent&&!e.showClose?(b(),oe(r,{key:0,class:x([e.nsSelect.e("caret"),e.nsSelect.e("icon"),e.iconReverse])},{default:le((()=>[(b(),oe(de(e.iconComponent)))])),_:1},8,["class"])):se("v-if",!0),e.showClose&&e.clearIcon?(b(),oe(r,{key:1,class:x([e.nsSelect.e("caret"),e.nsSelect.e("icon")]),onClick:e.handleClearClick},{default:le((()=>[(b(),oe(de(e.clearIcon)))])),_:1},8,["class","onClick"])):se("v-if",!0)])),_:2},[e.$slots.prefix?{name:"prefix",fn:le((()=>[C("div",ke,[y(e.$slots,"prefix")])]))}:void 0]),1032,["id","modelValue","placeholder","name","autocomplete","size","disabled","readonly","class","tabindex","onFocus","onBlur","onInput","onPaste","onCompositionstart","onCompositionupdate","onCompositionend","onKeydown"])],32)])),content:le((()=>[te(c,null,{default:le((()=>[m(te(d,{ref:"scrollbar",tag:"ul","wrap-class":e.nsSelect.be("dropdown","wrap"),"view-class":e.nsSelect.be("dropdown","list"),class:x([e.nsSelect.is("empty",!e.allowCreate&&Boolean(e.query)&&0===e.filteredOptionsCount)])},{default:le((()=>[e.showNewOption?(b(),oe(p,{key:0,value:e.query,created:!0},null,8,["value"])):se("v-if",!0),y(e.$slots,"default")])),_:3},8,["wrap-class","view-class","class"]),[[f,e.options.size>0&&!e.loading]]),e.emptyText&&(!e.allowCreate||e.loading||e.allowCreate&&0===e.options.size)?(b(),g(ae,{key:0},[e.$slots.empty?y(e.$slots,"empty",{key:0}):(b(),g("p",{key:1,class:x(e.nsSelect.be("dropdown","empty"))},S(e.emptyText),3))],64)):se("v-if",!0)])),_:3})])),_:3},8,["visible","placement","teleported","popper-class","effect","transition","persistent","onShow"])],34)),[[v,e.handleClose,e.popperPaneRef]])}],["__file","/home/runner/work/element-plus/element-plus/packages/components/select/src/select.vue"]]);var ze=h(r({name:"ElOptionGroup",componentName:"ElOptionGroup",props:{label:String,disabled:{type:Boolean,default:!1}},setup(e){const l=u("select"),a=w(!0),s=i(),r=w([]);Y(we,p({...d(e)}));const c=t(Le);L((()=>{r.value=v(s.subTree)}));const v=e=>{const t=[];return Array.isArray(e.children)&&e.children.forEach((e=>{var l;e.type&&"ElOption"===e.type.name&&e.component&&e.component.proxy?t.push(e.component.proxy):(null==(l=e.children)?void 0:l.length)&&t.push(...v(e))})),t},{groupQueryChange:h}=o(c);return n(h,(()=>{a.value=r.value.some((e=>!0===e.visible))}),{flush:"post"}),{visible:a,ns:l}}}),[["render",function(e,t,l,o,a,n){return m((b(),g("ul",{class:x(e.ns.be("group","wrap"))},[C("li",{class:x(e.ns.be("group","title"))},S(e.label),3),C("li",null,[C("ul",{class:x(e.ns.b("group"))},[y(e.$slots,"default")],2)])],2)),[[f,e.visible]])}],["__file","/home/runner/work/element-plus/element-plus/packages/components/select/src/option-group.vue"]]);const De=ce(Be,{Option:Ie,OptionGroup:ze}),qe=ve(Ie);ve(ze);export{qe as E,De as a};
