<script>
import {ref, onMounted, createCommentVNode} from "vue";
import Countdown from "../components/Countdown.vue";
import {CreateCountDown, GetCountDown} from "../../wailsjs/go/controller/App.js";
// import {useMessage} from "naive-ui";
import { useRouter, useRoute } from 'vue-router'
export default {
  name: "CountdownView",
  components:{ Countdown},
  setup() {
    let countDowns = ref([]);
    let showModal = ref(false);
    function init() {
      showModal.value = false ;
      formValue.value = {content:"",
        title:"",
        deadTime: 1704038400000,
        active: true,
        id : 0,
        duration: 0,};
      GetCountDown().then((res) => {
        console.log("getCountDown", res.data);
        countDowns.value = res.data;
      });
    }
    const createCountDown = () => {
      showModal.value = true;
      console.log(showModal.value);
    };
    let formValue = ref({
      content:"",
      title:"",
      deadTime: 1704038400000,  // millisecond
      active: true,
      id : 0,
      duration: 0,
    });
    onMounted(init);
    const formRef = ref(null);
    // const message = useMessage();
    const router = useRouter()
    const submit =() =>{
      console.log("submit");
      let arg1 = {
        title: formValue.value.title,
        content: formValue.value.content,
        deadTime: formValue.value.deadTime,
        active: formValue.value.active,
        id: 0,
        duration: 0,
      }
      CreateCountDown(arg1).then((res) => {
        console.log("res", res);
        if (res.code === 200){
          // message.success("创建成功");
          console.log("创建成功")
        }
        init();
      }).catch((err) => {
        console.log("err", err);
        // message.error("创建失败");
        console.log("arg1 ",arg1)
      });


    }
    return {
      countDowns,
      showModal,
      createCountDown,
      submit,
      formRef,
      title : ref("新建的表单"),
      size: ref("medium"),
      formValue,
      rules: {

      },
    };
  }
}

</script>

<template>
  <n-grid :x-gap="12" :y-gap="8" :cols="4">
    <n-grid-item v-for="(item,index) in countDowns">
      <Countdown class="my-countdown" :duration="item.duration" :active="item.active" :content="item.content" :id="item.id"
      :title="item.title"/>
    </n-grid-item>

  </n-grid>
  <n-p>
    <n-button type="success" @click="createCountDown">
      Add
    </n-button>
  </n-p>
  <n-modal v-model:show="showModal">
<!--    form  -->
    <n-card
        style="width: 600px; height: 400px; display: flex"
        :title="title"
        :bordered="false"
        size="huge"
        role="dialog"
        aria-modal="true"
    >

        <n-form
            ref="formRef"
            :label-width="80"
            :model="formValue"
            :rules="rules"
            size="medium"
            class="my-form"
        >
          <n-space>
            <n-form-item label="title" path="title">
              <n-input v-model:value="formValue.title" placeholder="title"/>
            </n-form-item>
            <n-form-item label="content" path="content">
              <n-input v-model:value="formValue.content" placeholder="content"/>
            </n-form-item>
          </n-space>

          <n-space>
            <n-form-item label="deadTime" path="deadTime">
              <n-date-picker v-model:value="formValue.deadTime" type="datetime" clearable />
            </n-form-item>

            <n-form-item label="active" path="active">
              <n-switch v-model:value="formValue.active" />
            </n-form-item>
          </n-space>

          <n-space>
<!--            <n-form-item>-->
<!--              <n-button attr-type="button" type="info" @click="">-->
<!--                验证-->
<!--              </n-button>-->
<!--            </n-form-item>-->
            <n-form-item>
              <n-button attr-type="submit" type="success" @click="submit">
                submit
              </n-button>
            </n-form-item>
          </n-space>
        </n-form>
    </n-card>
  </n-modal>
</template>

<style scoped>
.light-green {
  height: 108px;
  background-color: rgba(0, 128, 0, 0.12);
}
.green {
  height: 108px;
  background-color: rgba(0, 128, 0, 0.24);
}
.my-form{
  display: block;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  place-items: center;
  padding: 0 10%;
}
.my-countdown{
  height: 100%;
}
</style>