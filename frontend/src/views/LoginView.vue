<script>
import {useRouter} from "vue-router";
import {ref} from "vue";
import {CheckLogin} from "../../wailsjs/go/controller/App.js";

export default {
  name: "LoginView",
  setup(){
    const router = useRouter();
    const userinfo = ref({
      username : "",
      password : "",
    });
    const formRef = ref(null);
    const rules = {
      age :[
        {
          required: true,
          validator(rule,value){
            if (value === ""){
              return new Error("用户名不能为空");
            }
            return true;
          },
          trigger: ["input", "blur"]
        }
      ],
      password: [
        {
          required: true,
          message: "密码不能为空",
        }
      ]
    }
    const loginCheck= ()=>{
      console.log("login check");
      const userinfo2check = userinfo.value;
      console.log("userinfo2check", userinfo2check);
      CheckLogin(userinfo2check).then((res)=>{
        console.log("login check res", res)
        if(res.code === 200){
          console.log("login check success");
          const user = {
            id: res.data.id,
            description: res.data.description,
            avatar: res.data.avatar,
            email: res.data.email,
            username: res.data.username,
            password: res.data.password,
          };
          localStorage.setItem("user", JSON.stringify(user));
          router.push({path: "/home"});
        }else {
          console.log(res.msg);
        }
      }).catch((err)=>{
        console.log("login check failed", err);
      })
    };
    const register = ()=>{
      console.log("register");

    }
    return {
      userinfo,
      formRef,
      rules,
      loginCheck,
      register,
    }
  }
}
</script>

<template>
  <div class="my-login-form">
    <n-card title="LOGIN" hoverable style=" width: 400px">
      <n-form ref="formRef" :model="userinfo" :rules="rules"  >
        <n-form-item path="age" label="用户名">
          <n-input v-model:value="userinfo.username" @keydown.enter.prevent />
        </n-form-item>
        <n-form-item path="password" label="密码">
          <n-input
              v-model:value="userinfo.password"
              type="password"
              @keydown.enter.prevent
          />
        </n-form-item>
        <n-form-item>
          <div class="button-container">
            <n-space size="large">
              <n-button type="default">
                注册
              </n-button>
              <n-button @click="loginCheck" type="default">
                登陆
              </n-button>
            </n-space>
          </div>
        </n-form-item>

      </n-form>
    </n-card>
  </div>

</template>

<style scoped>
.my-login-form {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
.button-container {
  width : 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>