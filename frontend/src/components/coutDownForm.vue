<template>
  <n-modal v-model:show="showModal">
    <n-card
        style="width: 600px"
        :title="title"
        :bordered="false"
        size="huge"
        role="dialog"
        aria-modal="true"
    >
      <template #header-extra>
        顶部内容
      </template>
      <template class="my-form">
        <n-form
            ref="formRef"
            inline
            :label-width="80"
            :model="formValue"
            :rules="rules"
            size="medium"
        >
          <n-form-item label="title" path="title">
            <n-input v-model:value="formValue.title" placeholder="title"/>
          </n-form-item>
          <n-form-item label="content" path="content">
            <n-input v-model:value="formValue.content" placeholder="content"/>
          </n-form-item>

          <n-form-item label="deadTime" path="deadTime">
            <n-date-picker v-model:value="formValue.deadTime" type="datetime" clearable />
          </n-form-item>

          <n-form-item label="active" path="active">
            <n-switch v-model:value="formValue.active" />
          </n-form-item>

          <n-form-item>
            <n-button attr-type="button" @click="handleValidateClick">
              验证
            </n-button>
          </n-form-item>
          <n-form-item>
            <n-button attr-type="button" @click="submit">
              submit
            </n-button>
          </n-form-item>
        </n-form>

        <pre>{{ JSON.stringify(formValue, null, 2) }}</pre>
      </template>
      <template #footer>
        尾部
      </template>
    </n-card>
  </n-modal>
</template>

<script>
import { defineComponent, ref } from "vue";
// import {useMessage} from "naive-ui";
import {CreateCountDown} from "../../wailsjs/go/controller/App.js";
import { useRouter, useRoute } from 'vue-router'
export default defineComponent({
  props:{
    showModal: {
      type: Boolean,
      default : false
    },
    title:{
      type:String,
      default : "新建的表单"
    }
  },
  setup() {
    const formRef = ref(null);
    // const message = useMessage();
    const router = useRouter()
    const submit =() =>{
      console.log("submit");
      // sub mit the form
      // close the form
      // refresh the list
      CreateCountDown(this.formValue.value).then((res) => {
        console.log(res);
        if (res.data === 200){
          // message.success("创建成功");
        }
        router.push("/countDown/");
      });

    }
    return {
      formRef,
      submit,
      size: ref("medium"),
      formValue: ref({
        content:"",
        title:"",
        deadTime : 1672502400,
        active: true,
      }),
      rules: {
        user: {
          name: {
            required: true,
            message: "请输入姓名",
            trigger: "blur"
          },
          age: {
            required: true,
            message: "请输入年龄",
            trigger: ["input", "blur"]
          }
        },
        phone: {
          required: true,
          message: "请输入电话号码",
          trigger: ["input"]
        }
      },
      handleValidateClick(e) {
        e.preventDefault();
        formRef.value?.validate((errors) => {
          if (!errors) {
            // message.success("Valid");
          } else {
            console.log(errors);
            // message.error("Invalid");
          }
        });
      }
    };
  }
});
</script>