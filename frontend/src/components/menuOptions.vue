<template>
  <n-menu :options="menuOptions" />
</template>

<script>
import { defineComponent, h } from "vue";
import { RouterLink } from "vue-router";
import {
  LaptopOutline as WorkIcon,
  LogOutOutline as HomeIcon,
  ChatboxOutline as ChatIcon,
  TimerOutline as CountDownIcon,
  HeartCircle as HeartIcon,
} from "@vicons/ionicons5";
import {NIcon} from "naive-ui";
import chatView from "../views/ChatView.vue";

function renderIcon(icon) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const menuOptions = [
  {
    label: () => h(
        RouterLink,
        {
          to: {
            path: "/home",
          }
        },
        { default: () => "Home" }
    ),
    key: "go-back-home",
    icon: renderIcon(HeartIcon)
  },
  {
    label: () => h(
        RouterLink,
        {
          to: {
            path: "/about"
          }
        },
        { default: () => "上班" }
    ),
    key: "go-to-work",
    icon: renderIcon(WorkIcon)
  },
  {
    label: () => h(
        RouterLink,
        {
          to: {
            path: "/chat"
          }
        },
        { default: () => "聊天" },
    ),
    key: "go-to-chat",
    icon: renderIcon(ChatIcon)
  },
  {
    label: () => h(
        RouterLink,
        {
          to: {
            path: "/countdown"
          }
        },
        { default: () => "倒计时" },
    ),
    key: "go-to-countDown",
    icon: renderIcon(CountDownIcon)
  },
];

export default defineComponent({
  setup() {
    return {
      menuOptions
    };
  }
});
</script>