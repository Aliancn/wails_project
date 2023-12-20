<script setup>
import MenuOptions from "../components/menuOptions.vue";


</script>

<template>
  <div>
    <n-layout has-sider>
      <n-layout-sider bordered content-style="padding: 24px;">
        <menu-options></menu-options>
      </n-layout-sider>
      <n-layout>
        <n-layout-header bordered>
          这里准备增加一个标签的导航👋
        </n-layout-header>

        <n-layout-content content-style="padding: 24px;">
          <router-view></router-view>
        </n-layout-content>
<!--        <n-layout-footer bordered>-->
<!--          这里是页脚-->
<!--        </n-layout-footer>-->
      </n-layout>
    </n-layout>
  </div>
</template>

<style scoped>
.n-layout{
  height: 100vh;
}
.n-layout-header {
  padding: 24px;
}

.n-layout-footer {
  padding: 24px;
}
</style>