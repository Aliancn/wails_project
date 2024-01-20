<script>
import {useRouter} from "vue-router";
import {defineComponent, ref} from "vue";
import {CheckLogin} from "../../wailsjs/go/controller/App.js";
import {useMessage} from "naive-ui";

export default defineComponent({
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
    const message = useMessage();
    const showModal = ref(false);
    const regform = ref({
      username : "",
      password : "",
      email : "",
      avatar : "",
      description : "",
    });
    const loginCheck= ()=>{
      console.log("login check");
      const userinfo2check = userinfo.value;
      console.log("userinfo_to_check", userinfo2check);
      CheckLogin(userinfo2check).then((res)=>{
        console.log("login_check_res", res)
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
          message.error(res.msg);
        }
      }).catch((err)=>{
        console.log("login check failed", err);
      })
    };
    const register = ()=>{
      console.log("register");
      showModal.value = true;
    };
    // 经测试，localhost的存储是持久的
    // 就算程序关闭也会保存
    return {
      userinfo,
      formRef,
      rules,
      showModal,
      regform,
      loginCheck,
      register,
    }
  }
})
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
                <n-button type="default" @click="register">
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
  <n-modal v-model:show="showModal">
    <n-card
        style="width: 600px"
        title="注册"
        :bordered="false"
        size="huge"
        role="dialog"
        aria-modal="true"
    >
      <n-form ref="formRef" :model="regform" :rules="rules"  >
        <n-form-item path="age" label="用户名">
          <n-input v-model:value="regform.username" @keydown.enter.prevent />
        </n-form-item>
        <n-form-item path="password" label="密码">
          <n-input
              v-model:value="regform.password"
              type="password"
              @keydown.enter.prevent
          />
        </n-form-item>
        <n-form-item>
          <n-input>

          </n-input>
        </n-form-item>
      </n-form>
    </n-card>
  </n-modal>
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