<template>
  <div>
    <div>
      <div class="fullscreen-overlay" v-if="showOverlay">
        <div class="loading-wrapper">
          <div class="loading-spinner"></div>
          <div class="loading-text">视频处理中...</div>
        </div>
      </div>
    </div>
    <el-upload
        :before-upload="handleBeforeUpload"
        class="upload-demo"
        drag
        action="#"
        :http-request="uploadHttpRequest"
        multiple>
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      <div class="el-upload__tip" slot="tip">只能上传mp4文件</div>
    </el-upload>
  </div>
</template>

<script>
export default {
  name: 'UploadVideo',
  data() {
    return {
      customFile: '',
      showOverlay: false,
    }
  },
  methods: {
    handleBeforeUpload(file) {
      const fileType = file.type;
      const allowedTypes = ['video/mp4'];
      if (!allowedTypes.includes(fileType)) {
        this.$message.error('暂时只支持mp4 格式的文件');
        return false; // 阻止文件上传
      }
      this.showOverlay = true; // 在上传前显示遮罩
    },
    uploadHttpRequest(param) {
      this.customFile = param.file
      const data = new FormData()
      const fileUps = this.customFile
      data.append('file', fileUps)
      this.$axios.postForm('/upload-video', data, {
        headers: {'Content-Type': 'multipart/form-data'},
        responseType: 'blob'
      }).then(res => {
        let filename = res.headers['content-disposition'].split(';')[1].split('filename=')[1];
        var blob = new Blob([res.data]),
            Temp = document.createElement('a');
        Temp.href = window.URL.createObjectURL(blob);
        Temp.download = window.decodeURI(filename);
        document.body.appendChild(Temp);
        Temp.click();
        document.body.removeChild(Temp);
        window.URL.revokeObjectURL(Temp);
      })
    },
  }
}
</script>

<style scoped>
.fullscreen-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 9999;
  display: flex;
  justify-content: center;
  align-items: center;
}

.loading-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center; /* 添加文本居中的样式 */
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: 4px solid #fff;
  border-top-color: #888;
  animation: spin 1s infinite linear;
}

.loading-text {
  margin-top: 10px;
  color: #fff;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
</style>
