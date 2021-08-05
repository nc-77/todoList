<template>
  <el-container class="container">
    <el-header>
      <el-row justify="space-between">
        <el-col :span="8"> <h1>My Todo List</h1></el-col>
        <el-col :span="6"
          ><el-button type="success">Add Task</el-button>
        </el-col></el-row
      >
    </el-header>

    <el-main>
      <tasks
        :tasks="tasks"
        @delete-task="deleteTask"
        @update-Task="updateTask"
      />
    </el-main>

    <el-footer>
      <el-row justify="center">
        <p>Copyright &copy; 2021 NiC</p>
      </el-row>
    </el-footer>
  </el-container>
</template>

<script>
import Tasks from "../components/Tasks.vue";

let token =
  "Bear eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbiIsImV4cCI6MTYyNzY1MjIzMCwiaXNzIjoiTmlDIn0.fVr9ZWNBLUMlUHgQR5QsliJyPACpIaORVM9F3Y1nIrM";

let ErrUnknown = "出错了,请联系管理员";
let ErrToken = "认证信息过期,请重新登录";

export default {
  name: "Home",
  components: {
    Tasks,
  },
  data() {
    return {
      tasks: [],
      showAddTask: false,
    };
  },
  async created() {
    this.tasks = await this.fetchData();
  },

  methods: {
    async addTask(newTask) {
      const res = await fetch("api/tasks", {
        method: "PUT",
        headers: {
          Authorization: token,
        },
        body: JSON.stringify(newTask),
      });
      const result = await res.json();
      this.refresh(result.code);
    },

    async deleteTask(id) {
      const res = await fetch("api/tasks", {
        method: "DELETE",
        headers: {
          Authorization: token,
        },
        body: JSON.stringify({ id: [id] }),
      });
      const result = await res.json();
      this.refresh(result.code);
    },

    async updateTask(newTask) {
      const res = await fetch("api/tasks", {
        method: "POST",
        headers: {
          Authorization: token,
        },
        body: JSON.stringify(newTask),
      });

      const result = await res.json();
      this.refresh(result.code);
    },

    toggleAddTask() {
      this.showAddTask = !this.showAddTask;
    },
    async fetchData() {
      const res = await fetch("api/tasks", {
        headers: {
          Authorization: token,
        },
      });
      const resData = await res.json();
      if (resData.code !== 200) {
        // 错误处理
        switch (resData.code) {
          case 401:
            alert(ErrToken);
            break;
          default:
            alert(ErrUnknown);
            break;
        }
        return;
      }
      const tasks = resData.data;
      tasks.sort(function (a, b) {
        if (a.isFinish === b.isFinish) {
          if (a.isFinish) {
            return a.ddl < b.ddl ? 1 : -1;
          } else {
            return a.ddl > b.ddl ? 1 : -1;
          }
        } else {
          return a.isFinish ? 1 : -1;
        }
      });

      return tasks;
    },
    async refresh(httpCode) {
      switch (httpCode) {
        case 200:
          this.tasks = await this.fetchData();
          break;
        case 401:
          alert(ErrToken);
          // 跳转至登录页 todo

          break;
        default:
          alert(ErrUnknown);
      }
    },
  },
};
</script>

<style scoped>
.container {
  max-width: 500px;
  margin: 30px auto;
  overflow: auto;
  min-height: 300px;
  border: 1px solid steelblue;
  padding: 30px 30px 10px 30px;
  border-radius: 5px;
}
.el-header {
  padding: 0;
}
.el-main {
  padding: 0;
}
.el-footer {
  height: 20px;
  font-size: 8px;
}
</style>