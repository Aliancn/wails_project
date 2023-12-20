
<template>
  <n-card :title="title" hoverable content-style="padding: 3px;">
    <n-space class="my-space" vertical>
      <span>
        {{content}}
      </span>

      <span style="font-variant-numeric: tabular-nums">
        <n-countdown :duration="duration" :active="active" />
      </span>
<!--      <span>-->
<!--          <n-switch v-model:value="active" />-->
<!--      </span>-->
      <span>
        <n-button type="error" @click="del">
          删除
        </n-button>
      </span>
    </n-space>
  </n-card>
</template>

<script>
import {defineComponent, onMounted, ref} from "vue";
import {DeleteCountDown} from "../../wailsjs/go/controller/App.js";

export default defineComponent({
  props:{
    duration: {
      type: Number,
      default: 0
    },
    active: {
      type: Boolean,
      default: true
    },
    content: {
      type: String,
      default: "倒计时"
    },
    title: {
      type: String,
      default: "倒计时"
    },
    id: {
      type: Number,
      default: "id"
    }
  },
  setup(props) {
    const del = () => {
      console.log("delete" , props.id)
      const _id =props.id;
      DeleteCountDown(_id).then((res) => {
        console.log("res", res);
        if (res.code === 200){
          // message.success("创建成功");
          console.log("删除成功")
          location.reload();
        }
      }).catch((err) => {
        console.log("err", err);
      })
    };
    return {
      del,

    };
  }
});
</script>

<style scoped>
.n-card {
  max-width: 300px;
  margin: 0 auto;
}
.my-space{
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  place-items: center;
}

span{
  margin: 0 8px;
  width: 100%;
}
</style>