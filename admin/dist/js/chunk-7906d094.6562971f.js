(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-7906d094"],{"0bff":function(t,s,o){"use strict";o.r(s);var a=function(){var t=this,s=t.$createElement,o=t._self._c||s;return o("section",{staticClass:"auth reset-password"},[o("div",{staticClass:"container"},[o("div",{staticClass:"row auth-row"},[o("div",{staticClass:"col-12"},[t.getLogo?o("figure",{staticClass:"auth-logo"},[o("img",{attrs:{src:t.globalBasePath+t.getLogo}})]):t._e(),o("div",{staticClass:"auth-card"},[o("div",{staticClass:"auth-card-cont"},[t._m(0),o("form",{staticClass:"form form-center"},[o("div",{staticClass:"form-group"},[o("input",{directives:[{name:"model",rawName:"v-model",value:t.password,expression:"password"}],staticClass:"form-input",attrs:{type:"password",autocomplete:"new-password",placeholder:"Password"},domProps:{value:t.password},on:{input:function(s){s.target.composing||(t.password=s.target.value)}}})]),o("router-link",{staticClass:"auth-link",attrs:{to:{name:"login"}}},[t._v("Back to login")]),o("div",{staticClass:"auth-btn-cont"},[o("button",{staticClass:"btn btn-arrow btn-transparent btn-arrow",attrs:{type:"submit"},on:{click:function(s){return s.preventDefault(),t.doReset(s)}}},[t._v("Reset")])])],1)])])])])])])},e=[function(){var t=this,s=t.$createElement,o=t._self._c||s;return o("div",{staticClass:"auth-text"},[o("h2",[t._v("Forgot password?")]),o("p",[t._v("Enter something about password here")])])}],n={name:"ResetPassword",data:function(){return{doingAxios:!1,password:"",token:""}},mounted:function(){this.token=this.$route.params.token,this.doVerify()},methods:{doVerify:function(){this.axios.get("/email/verify/"+this.token)},doReset:function(){this.axios.post("/password/reset",{password:this.password,token:this.token}).then((function(t){console.log(t)})).catch((function(t){console.log(t)}))}},computed:{getLogo:function(){return this.$store.state.logo}}},r=n,i=(o("d3d7"),o("2877")),c=Object(i["a"])(r,a,e,!1,null,"2158ea28",null);s["default"]=c.exports},"8fc4":function(t,s,o){},d3d7:function(t,s,o){"use strict";var a=o("8fc4"),e=o.n(a);e.a}}]);
//# sourceMappingURL=chunk-7906d094.6562971f.js.map