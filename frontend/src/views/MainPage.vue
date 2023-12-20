<script>
import chtholly from '../assets/images/chtholly.jpg'
import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import routes from "../router/routes.js";
export default {
  name: "MainPage",
  setup(){
    const router = useRouter();
    const user = ref({
      id: 1,
      description: "这是一个测试用户",
      avatar: 'src/assets/images/chtholly.jpg',
      email: "180@qq.com",
      username: "aliancn",
    })
    const init = ()=>{
      console.log("init main page");
      if (localStorage.getItem("user") === null) {
        // TODO: 这里应该是一个跳转到登录页面的操作
        router.push('/login/');
      } else {
        user.value = JSON.parse(localStorage.getItem("user"));
        console.log("user", user.value);
      }
    };
    const checkout = ()=>{
      localStorage.removeItem("user");
      console.log("checkout")
      router.push('/login/');
    }
    onMounted(()=>{
      init();
    });
    return {
      user,
      chtholly,
      checkout,
    }
  }
}
</script>

<template>
  <n-layout>
    <n-layout-header>
      <n-avatar
          size="large"
          round
          :src="user.avatar"
      />
    </n-layout-header>
    <n-layout-content content-style="padding: 24px;" class="my-page-main">
      <n-list hoverable clickable>
        <n-list-item>
          <n-thing title="ID" content-style="margin-top: 10px;">
            {{ user.id }}
          </n-thing>
        </n-list-item>
        <n-list-item>
          <n-thing title="Description" content-style="margin-top: 10px;">
            {{ user.description}}
          </n-thing>
        </n-list-item>
      </n-list>
    </n-layout-content>
    <n-layout-footer >
      <n-space size="large" vertical>
        <n-button type="primary" >修改</n-button>
        <n-button type="primary"  @click="checkout">退出</n-button>
      </n-space>
    </n-layout-footer>
  </n-layout>
</template>

<style scoped>
.my-page-main {
  height: 70vh;
}
</style>