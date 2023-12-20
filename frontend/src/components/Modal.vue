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
        噢！
      </template>
      内容
      <template #footer>
        尾部
      </template>
    </n-card>
  </n-modal>
</template>

<script>
import { defineComponent, ref } from "vue";

export default defineComponent({
  props:{
    showModal: {
      type: Boolean,
      default : false
    },
    title:{
      type:String,
      default : "模态框"
    }
  },
  setup() {
    return {
    };
  }
});
</script>