/*! 
 Build based on gin-vue-admin 
 Time : 1679909256000 */
import{a7 as n,m as t,a8 as e,a9 as r,aa as i}from"./index.61dd0a2d.js";var a=/\s/;var o=/^\s+/;function u(n){return n?n.slice(0,function(n){for(var t=n.length;t--&&a.test(n.charAt(t)););return t}(n)+1).replace(o,""):n}var s=NaN,c=/^[-+]0x[0-9a-f]+$/i,f=/^0b[01]+$/i,d=/^0o[0-7]+$/i,l=parseInt;function v(e){if("number"==typeof e)return e;if(n(e))return s;if(t(e)){var r="function"==typeof e.valueOf?e.valueOf():e;e=t(r)?r+"":r}if("string"!=typeof e)return 0===e?e:+e;e=u(e);var i=f.test(e);return i||d.test(e)?l(e.slice(2),i?2:8):c.test(e)?s:+e}const m=function(){return e.Date.now()};var p="Expected a function",g=Math.max,h=Math.min;function x(n,e,r){var i,a,o,u,s,c,f=0,d=!1,l=!1,x=!0;if("function"!=typeof n)throw new TypeError(p);function y(t){var e=i,r=a;return i=a=void 0,f=t,u=n.apply(r,e)}function T(n){var t=n-c;return void 0===c||t>=e||t<0||l&&n-f>=o}function b(){var n=m();if(T(n))return w(n);s=setTimeout(b,function(n){var t=e-(n-c);return l?h(t,o-(n-f)):t}(n))}function w(n){return s=void 0,x&&i?y(n):(i=a=void 0,u)}function E(){var n=m(),t=T(n);if(i=arguments,a=this,c=n,t){if(void 0===s)return function(n){return f=n,s=setTimeout(b,e),d?y(n):u}(c);if(l)return clearTimeout(s),s=setTimeout(b,e),y(c)}return void 0===s&&(s=setTimeout(b,e)),u}return e=v(e)||0,t(r)&&(d=!!r.leading,o=(l="maxWait"in r)?g(v(r.maxWait)||0,e):o,x="trailing"in r?!!r.trailing:x),E.cancel=function(){void 0!==s&&clearTimeout(s),f=0,i=c=a=s=void 0},E.flush=function(){return void 0===s?u:w(m())},E}const y=new Map;let T;function b(n,t){let e=[];return Array.isArray(t.arg)?e=t.arg:i(t.arg)&&e.push(t.arg),function(r,i){const a=t.instance.popperRef,o=r.target,u=null==i?void 0:i.target,s=!t||!t.instance,c=!o||!u,f=n.contains(o)||n.contains(u),d=n===o,l=e.length&&e.some((n=>null==n?void 0:n.contains(o)))||e.length&&e.includes(u),v=a&&(a.contains(o)||a.contains(u));s||c||f||d||l||v||t.value(r,i)}}r&&(document.addEventListener("mousedown",(n=>T=n)),document.addEventListener("mouseup",(n=>{for(const t of y.values())for(const{documentHandler:e}of t)e(n,T)})));const w={beforeMount(n,t){y.has(n)||y.set(n,[]),y.get(n).push({documentHandler:b(n,t),bindingFn:t.value})},updated(n,t){y.has(n)||y.set(n,[]);const e=y.get(n),r=e.findIndex((n=>n.bindingFn===t.oldValue)),i={documentHandler:b(n,t),bindingFn:t.value};r>=0?e.splice(r,1,i):e.push(i)},unmounted(n){y.delete(n)}};export{w as C,x as d};
