<script>

import {ref} from "vue";
import {CreateTodoList, DeleteTodoList, GetTodoList, UpdateTodoList} from "../../wailsjs/go/controller/App.js";

export default{
  name: "TodoListView",
  data(){
    return {
      headerTabs :[
        { value:"allItem", name:"全部" },
        { value:"processing", name:"進行中" },
        { value:"done", name:"已完成" }
      ],
      visibilityTab : "",
      todoList : [
        {
          id:  "201904221312",
          taskName: '進行中的任務',
          completed: false,
        },
        {
          id:  "201904221313",
          taskName: '已完成的任務',
          completed: true,
        },
      ],
      newTask: "",
      cacheTask: {},
      cacheTaskName: "",
      isEdit: false
    }
  },
  created(){
    this.visibilityTab = this.headerTabs[0].value ;
    GetTodoList().then((res) => {
      console.log("getTodoList", res.data);
      this.todoList = res.data;
    });
  },
  methods:{
    addToList(){
      let timestamp =  Math.floor(Date.now());
      let name = this.newTask.trim();
      if (name){
        // this.todoList.push({
        //   id: timestamp,
        //   taskName:  name,
        //   completed: false
        // })
        let arg1 = {
          id: timestamp,
          taskName: name,
          completed: false
        }
        CreateTodoList(arg1).then((res) => {
          console.log("res", res);
          if (res.code === 200){
            console.log("创建成功")
          }
          GetTodoList().then((res) => {
            console.log("getTodoList", res.data);
            this.todoList = res.data;
          });
        });
      }
      this.newTask = "";
    },
    deleteTask(task){
      // let removeId = this.todoList.findIndex((item)=>{
      //   return item.id === task.id
      // })
      // this.todoList.splice(removeId, 1);
      let arg1 = {
        id: task.id,
      }
      DeleteTodoList(arg1).then((res) => {
        console.log("res", res);
        if (res.code === 200){
          console.log("删除成功")
        }
        GetTodoList().then((res) => {
          console.log("getTodoList", res.data);
          this.todoList = res.data;
        });
      });
    },
    deleteAllTask(){
      this.todoList = [];
    },
    editTask(item){
      this.isEdit = true;
      this.cacheTask = item;
      this.cacheTaskName = item.taskName;
    },
    cancelEdit(){
      this.isEdit = false;
      this.cacheTask = {};
    },
    doneEdit(item){
      this.isEdit = false;
      item.taskName = this.cacheTaskName;
      this.cacheTask = {};
      this.cacheTaskName = "";
      let arg = {
        id: item.id,
        taskName: item.taskName,
        completed: item.completed
      }
      UpdateTodoList(arg).then((res) => {
        console.log("res", res);
        if (res.code === 200){
          console.log("更新成功")
        }
      });
    },
    lostFocus(item){
      if(this.isEdit){
        this.doneEdit(item);
      }
    }
  },
  computed: {
    filteredList(){
      let allTab =  this.headerTabs[0].value  ;
      let processingTab =  this.headerTabs[1].value  ;

      if (this.visibilityTab === allTab){
        return this.todoList;
      } else if (this.visibilityTab ===  processingTab){
        return this.todoList.filter((item)=>{
          return !item.completed
        });
      } else {
        return this.todoList.filter((item) => {
          return item.completed
        });
      }
    },
    uncompletedPrompt() {
      let message = "列表上沒有任務，添加一點咩～";
      let counter = 0 ;

      if (this.todoList.length <= 0){
        return message
      }

      this.todoList.forEach((item) => {
        counter =  !item.completed? counter +1 : counter ;
      });

      if (counter > 0){
        message = "還有 "+ counter + " 筆任務未完成，別偷懶了！";
      } else {
        message = "任務都完成了! 好棒 XD"
      }

      return message
    }
  }
}
</script>

<template>
  <div id="app" class="container my-3">
    <div class="input-group mb-3">
      <div class="input-group-prepend">
        <span class="input-group-text" id="basic-addon1">待辦事項</span>
      </div>
      <input type="text" class="form-control" placeholder="準備要做的任務"
             v-model="newTask"
             @keyup.enter="addToList()">
      <div class="input-group-append">
        <button class="btn btn-primary" type="button" @click="addToList()">新增</button>
      </div>
    </div>
    <div class="card text-center">
      <div class="card-header">
        <ul class="nav nav-tabs card-header-tabs">
          <li class="nav-item" v-for="tab in headerTabs">
            <a class="nav-link" href="#"
               :class="{'active': visibilityTab === tab.value}"
               @click="visibilityTab = tab.value">{{tab.name}}</a>
          </li>
        </ul>
      </div>
      <ul class="list-group list-group-flush text-left">
        <li class="list-group-item" v-for="(item, index) in filteredList" >
          <div class="d-flex"  @dblclick="editTask(item)" v-if="item.id !== cacheTask.id" >
            <div class="form-check">
              <input type="checkbox" class="form-check-input" :id="item.id"
                     v-model="item.completed">
              <label class="form-check-label"
                     :for="item.id"
                     :class="{'completed':item.completed}">
                {{item.taskName}}
              </label>
            </div>
            <button type="button" class="close ml-auto" aria-label="Close" @click="deleteTask(item)">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <input type="text"
                 class="form-control"
                 autofocus
                 v-else
                 v-model="cacheTaskName"
                 @keyup.esc="cancelEdit()"
                 @keyup.enter="doneEdit(item)"
                 @blur="lostFocus(item)">
        </li>
      </ul>
      <div class="card-footer d-flex justify-content-between">
        <span>{{uncompletedPrompt}}</span>
        <a href="#" @click.prevent="deleteAllTask()">清除所有任務</a>
      </div>
    </div>
  </div>
</template>

<style scoped>
.completed {
  text-decoration: line-through;
}

</style>

<style scoped>
@import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.8.1/css/all.min.css');
@import url('https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.0.0/css/bootstrap.css');

</style>
