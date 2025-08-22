<template>
  <div class="setting-container">
    <div class="setting-header fade-in-down">
      <div class="setting-title">设置</div>
    </div>

    <div class="setting-wrapper fade-in-up">
      <div class="settings-section delay-1">
        <h3 class="section-title">外观设置</h3>
        <div class="setting-item">
          <label class="setting-label">主题</label>
          <div class="setting-control">
            <DropDown :list="themeList" :placeholder="placeholder" v-model:value="ThemeValue"/>
          </div>
        </div>
      </div>

      <div class="settings-section delay-2">
        <h3 class="section-title">QQ机器人设置</h3>
        <div class="setting-item"
             style="display: grid; grid-template-columns: 1fr 1fr; padding: 10px; gap: 10px">

          <div style="display: flex; align-items: center">
            <label class="setting-label">机器人开关</label>
            <div class="setting-control">
              <CustomSwitch
                  v-model="LLOneBotSwitch"
              />
            </div>
          </div>

          <div style="display: flex; align-items: center">
            <label class="setting-label">LLOneBot http端口</label>
            <div class="setting-control">
              <CustomInput
                  type="number"
                  :min="1"
                  :max="65535"
                  v-model="OneBotPort"
                  placeholder="请输入端口号"
              />
            </div>
          </div>

          <div style="display: flex; align-items: center">
            <label class="setting-label">LLOneBot Token</label>
            <div class="setting-control">
              <CustomInput
                  type="text"
                  placeholder="请输入Token"
                  v-model="LLOneBotToken"
              />
            </div>
          </div>

          <div style="display: flex; align-items: center">
            <label class="setting-label">推送群号</label>
            <div class="setting-control">
              <CustomInput
                  type="text"
                  placeholder="请输入群号"
                  v-model="QQGroup"
              />
            </div>
          </div>

        </div>
      </div>

      <div class="settings-section delay-2">
        <h3 class="section-title">服务器设置</h3>

        <div class="setting-item">
          <label class="setting-label">MC服务器目录</label>
          <div class="setting-control directory-control">
            <input
                type="text"
                class="directory-input"
                v-model="mcServerDirPath"
                placeholder="请选择MC服务器目录"
                readonly
            />
            <button class="browse-btn" @click="setMcServerDirPath()">
              <IconToggle class="browse-btn-icon" :FillIcon="IconFileUploadFill" :LineIcon="IconFileUploadLine"
                          :Size="20"
                          :Toggle="true"></IconToggle>
              选择
            </button>
            <button class="browse-btn" @click="openDirectory(mcServerDirPath!)">
              <IconToggle class="browse-btn-icon" :FillIcon="IconFolderOpenFill" :LineIcon="IconFolderOpenLine"
                          :Size="20"
                          :Toggle="true"></IconToggle>
              打开
            </button>
          </div>
        </div>

        <div class="setting-item">
          <label class="setting-label">代理服务器目录</label>
          <div class="setting-control directory-control">
            <input
                type="text"
                class="directory-input"
                v-model="proxyServerDirPath"
                placeholder="请选择代理服务器目录"
                readonly
            />
            <button class="browse-btn" @click="setProxyServerDirPath()">
              <IconToggle class="browse-btn-icon" :FillIcon="IconFileUploadFill" :LineIcon="IconFileUploadLine"
                          :Size="20"
                          :Toggle="true"></IconToggle>
              选择
            </button>
            <button class="browse-btn" @click="openDirectory(proxyServerDirPath!)">
              <IconToggle class="browse-btn-icon" :FillIcon="IconFolderOpenFill" :LineIcon="IconFolderOpenLine"
                          :Size="20"
                          :Toggle="true"></IconToggle>
              打开
            </button>
          </div>
        </div>
      </div>

      <div class="settings-section delay-3">
        <h3 class="section-title">关于</h3>
        <div class="about-content">
          <p>Voxesis Minecraft 服务器管理面板</p>
          <p>版本: v0.0.5-beta</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import DropDown from "../components/DropDown.vue"
import CustomInput from "../components/CustomInput.vue";
import IconToggle from "../components/IconToggle.vue";
import CustomSwitch from "../components/CustomSwitch.vue";
import {ElMessage} from "element-plus"
import {onMounted, ref, watch} from "vue";
import {
  getAllAppConfig,
  openMcServerDirectoryDialog,
  openProxyServerDirectoryDialog,
  upDataAppConfig
} from "../utils/SystemMonitor";
import {IconFileUploadFill, IconFileUploadLine, IconFolderOpenFill, IconFolderOpenLine} from "birdpaper-icon";
import {BrowserOpenURL} from '../../wailsjs/runtime';

type ThemeName = '亮色' | '暗色' | '森林' | '海洋' | '樱花' | '朋克';

type ThemeValueType = 'dark-theme' | 'forest-theme' | 'ocean-theme' |
    'sakura-theme' | 'synthwave-theme' | '';

const MOUNTED = ref<boolean>(false);

const themeList = ref<{ label: ThemeName; value: ThemeValueType }[]>([
  {label: '亮色', value: ''},
  {label: '暗色', value: 'dark-theme'},
  {label: '森林', value: 'forest-theme'},
  {label: '海洋', value: 'ocean-theme'},
  {label: '樱花', value: 'sakura-theme'},
  {label: '朋克', value: 'synthwave-theme'}
])

const ThemeValue = ref<ThemeValueType>();
const placeholder = ref<string>()

const OneBotPort = ref<string | number>()
const LLOneBotSwitch = ref<boolean>(false)
const LLOneBotToken = ref<string>()
const QQGroup = ref<string>()

const mcServerDirPath = ref<string>()
const proxyServerDirPath = ref<string>()

function setTheme(theme: ThemeValueType) {
  if (!MOUNTED.value) return;

  document.body.classList.remove('dark-theme', 'forest-theme', 'ocean-theme', 'sakura-theme', 'synthwave-theme');
  upDataAppConfig('theme', theme).then(() => {
    ElMessage({
      message: '主题修改成功',
      type: 'success',
    })
  });

  if (theme) document.body.classList.add(theme);
}

function setMcServerDirPath() {
  if (!MOUNTED.value) return;

  openMcServerDirectoryDialog().then(path => {
    if (path == "") {
      return ElMessage({
        message: '操作取消',
        type: 'warning',
      })
    }

    upDataAppConfig('mc_server_root_path', path).then(() => {
      mcServerDirPath.value = path;
      ElMessage({
        message: '设置成功',
        type: 'success',
      })
    })
  })
}

function setProxyServerDirPath() {
  if (!MOUNTED.value) return;

  openProxyServerDirectoryDialog().then(path => {
    if (path == "") {
      return ElMessage({
        message: '操作取消',
        type: 'warning',
      })
    }

    upDataAppConfig('mc_server_proxy_root_path', path).then(() => {
      proxyServerDirPath.value = path;
      ElMessage({
        message: '设置成功',
        type: 'success',
      })
    })
  })
}

function openDirectory(path: string) {
  if (!MOUNTED.value) return;

  if (path == "" || !path) {
    return ElMessage({
      message: '未设置目录',
      type: 'warning',
    })
  }

  BrowserOpenURL(`file://${path}`)
}

watch(ThemeValue, () => {
  if (!MOUNTED.value) return;

  setTheme(ThemeValue.value || "");
})

watch(OneBotPort, async () => {
  if (!MOUNTED.value) return;

  upDataAppConfig('qq_bot_port', OneBotPort.value).then(() => {
    ElMessage({
      message: '设置成功',
      type: 'success',
    })
  })
})

watch(LLOneBotSwitch, () => {
  if (!MOUNTED.value) return;

  upDataAppConfig('qq_bot', LLOneBotSwitch.value).then(() => {
    ElMessage({
      message: '设置成功',
      type: 'success',
    })
  })
})

watch(LLOneBotToken, () => {
  if (!MOUNTED.value) return;

  upDataAppConfig('llonebot_token', LLOneBotToken.value).then(() => {
    ElMessage({
      message: '设置成功',
      type: 'success',
    })
  })
})

watch(QQGroup, () => {
  if (!MOUNTED.value) return;

  upDataAppConfig('qq_group', QQGroup.value).then(() => {
    ElMessage({
      message: '设置成功',
      type: 'success',
    })
  })
})

onMounted(() => {
  getAllAppConfig().then(conf => {
    OneBotPort.value = conf.qq_bot_port;
    mcServerDirPath.value = conf.mc_server_root_path || '未设置';
    proxyServerDirPath.value = conf.mc_server_proxy_root_path || '未设置';
    LLOneBotSwitch.value = conf.qq_bot;
    LLOneBotToken.value = conf.llonebot_token;
    QQGroup.value = conf.qq_group;

    for (const item of themeList.value) {
      if (item.value == conf.theme) {
        placeholder.value = item.label;
        ThemeValue.value = item.value;
      }
    }

    setTimeout(() => MOUNTED.value = true)
  })
})
</script>

<style scoped>
.setting-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  background-color: var(--color-background-app);
  padding: 24px;
  box-sizing: border-box;
  overflow-y: auto;
}

.setting-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid var(--color-border-default);
}

.setting-title {
  font-size: 24px;
  font-weight: bold;
  color: var(--color-text-primary);
}

.setting-wrapper {
  position: relative;
  flex: 1;
}

.settings-section {
  background-color: var(--color-background-card);
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.section-title {
  font-size: 18px;
  font-weight: 600;
  color: var(--color-text-header);
  margin: 0 0 16px 0;
  padding-bottom: 8px;
  border-bottom: 1px solid var(--color-border-default);
}

.setting-item {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}

.setting-item:last-child {
  margin-bottom: 0;
}

.setting-label {
  width: 150px;
  font-size: 14px;
  color: var(--color-text-primary);
  font-weight: 500;
}

.setting-control {
  flex: 1;
}

.directory-control {
  display: flex;
  gap: 10px;
}

.directory-input {
  flex: 1;
  padding: 8px 12px;
  border: 1px solid var(--color-border-default);
  border-radius: 4px;
  background-color: var(--color-background-card);
  color: var(--color-text-primary);
  font-size: 14px;
}

.directory-input:focus {
  outline: none;
  border-color: var(--color-border-focus);
}

.browse-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 12px;
  background-color: var(--color-background-menu);
  color: var(--color-text-primary);
  border: 1px solid var(--color-border-default);
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  transition: background-color 0.2s;
}

.browse-btn:hover {
  background-color: var(--color-background-menu-hover);
}

.browse-btn-icon {
  transform: translateX(-10px);
}

.about-content {
  color: var(--color-text-secondary);
  font-size: 14px;
}

.about-content p {
  margin: 4px 0;
}

.fade-in-down {
  animation: fadeInDown 0.6s ease-out forwards;
  opacity: 0;
  transform: translateY(-20px);
}

.fade-in-up {
  animation: fadeInUp 0.6s ease-out forwards;
  opacity: 0;
  transform: translateY(20px);
}

.delay-1 {
  animation-delay: 0.1s;
}

.delay-2 {
  animation-delay: 0.2s;
}

.delay-3 {
  animation-delay: 0.3s;
}

@keyframes fadeInDown {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeInUp {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (max-width: 768px) {
  .setting-container {
    padding: 16px;
  }

  .setting-item {
    flex-direction: column;
    align-items: flex-start;
  }

  .setting-label {
    width: auto;
    margin-bottom: 8px;
  }

  .directory-control {
    width: 100%;
  }
}
</style>
