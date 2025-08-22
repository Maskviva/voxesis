<template>
  <header class="header" @dblclick="WinMaxSize = !WinMaxSize; winMaximize(WinMaxSize)">
    <a v-top-text v-hover v-ripple class="header-btn" style="padding: 8px;">
      <span style="display: flex; gap: 5px"><img style="margin: 0" src="./assets/images/logo_no_background.png"
                                                 width="25" height="25" alt="logo">Voxesis</span>
    </a>

    <div class="header-control">
      <a v-top-text v-hover class="header-btn" @click="winMinimize()">
        <IconSubtractLine/>
      </a>
      <a v-top-text v-hover class="header-btn" @click="WinMaxSize = !WinMaxSize; winMaximize(WinMaxSize)">
        <component :is="!WinMaxSize ? IconCheckboxBlankLine : IconCheckboxMultipleBlankLine"/>
      </a>
      <a v-top-text v-hover class="header-btn" @click="closeWin()">
        <IconCloseLargeLine/>
      </a>
    </div>
  </header>

  <ul ref="view_list_box" class="sidebar paternal-box">
    <li v-for="item in VIEW_LIST" @click="toggle_view(item.name)">
      <IconToggle :LineIcon="item.line_icon" :FillIcon="item.fill_icon" :Size="25"
                  :Toggle="view_component.name == item.name"></IconToggle>
      <span class="sidebar-text">{{ item.introduce }}</span>
    </li>
  </ul>

  <div class="view-box">
    <transition name="fade">
      <keep-alive include="dashboard">
        <component :is="view_component.component"/>
      </keep-alive>
    </transition>
  </div>
</template>

<script setup lang="ts">
import {onMounted, provide, ref, shallowRef} from 'vue';
import {
  IconCheckboxBlankLine,
  IconCheckboxMultipleBlankLine,
  IconCloseLargeLine,
  IconDashboard2Fill,
  IconDashboard2Line,
  IconHome3Fill,
  IconHome3Line,
  IconSettingsFill,
  IconSettingsLine,
  IconSubtractLine,
  IconTerminalBoxFill,
  IconTerminalBoxLine,
} from 'birdpaper-icon'

import {vRipple} from './utils/waves'
import {vHover} from './utils/hover'
import {vTopText} from './utils/topText'
import {closeWin, getAppConfigByKey, watchWindowState, winMaximize, winMinimize,} from './utils/SystemMonitor'
import IconToggle from "./components/IconToggle.vue";

import HomeView from "./view/home.vue"
import DashboardView from "./view/dashboard.vue"
import TerminalView from "./view/terminal.vue"
import SettingView from "./view/setting.vue"

const VIEW_LIST = [
  {
    name: 'home',
    component: HomeView,
    introduce: "欢迎页",
    line_icon: IconHome3Line,
    fill_icon: IconHome3Fill,
  },
  {
    name: 'dashboard',
    component: DashboardView,
    introduce: "仪表盘",
    line_icon: IconDashboard2Line,
    fill_icon: IconDashboard2Fill,
  },
  {
    name: 'terminal',
    component: TerminalView,
    introduce: "控制台",
    line_icon: IconTerminalBoxLine,
    fill_icon: IconTerminalBoxFill,
  },
  {
    name: 'setting',
    component: SettingView,
    introduce: "设置",
    line_icon: IconSettingsLine,
    fill_icon: IconSettingsFill,
  },
]

const WinMaxSize = ref(false);
const view_component = shallowRef({
  name: VIEW_LIST[0].name,
  component: VIEW_LIST[0].component
});
const sidebar_before_top = ref()
const view_list_box = ref()

const toggle_view = (view: typeof VIEW_LIST[number]['name']) => {
  const targetView = VIEW_LIST.find(item => item.name === view);

  sidebar_before_top.value = VIEW_LIST.indexOf(targetView!) * 40 + 10 + "px";

  if (targetView) {
    view_component.value = {
      name: view,
      component: targetView.component
    };
  }
}

onMounted(() => {
  getAppConfigByKey('theme').then((theme) => {
    document.body.classList.add(theme)
  })

  view_list_box.value.addEventListener('wheel', (event: WheelEvent) => {
    const ScrollDirection = event.deltaY > 0;
    const ViewIndex = VIEW_LIST.findIndex(component => component.name === view_component.value.name);
    if (ScrollDirection) {
      toggle_view(VIEW_LIST[ViewIndex + 1].name)
    } else {
      toggle_view(VIEW_LIST[ViewIndex - 1].name)
    }
  })

  watchWindowState((isMaximized) => {
    WinMaxSize.value = isMaximized;
  });
});

provide('AppViewMethod', {
  toggle_view,
})
</script>

<style scoped>
header {
  width: 100%;
  height: 50px;
  padding: 10px;
  background-color: var(--color-background-header);
  border-bottom: 1px solid var(--color-border-default);
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);

  --wails-draggable: drag;

  display: flex;
  align-items: center;

  z-index: 99999;

  transition: all 0.3s ease;
}

.header-control {
  margin-left: auto;
  padding: 0;
  --wails-draggable: no-drag;
  display: flex;
  flex-direction: row;
}

.header-btn {
  color: var(--color-text-header);
  text-decoration: none;
  padding: 8px 15px;
  margin: 10px;
  border-radius: 8px;
  cursor: pointer;
  --wails-draggable: no-drag;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  z-index: 1;

  transition: all 0.2s ease;
}

.header-btn:hover {
  background-color: var(--color-background-header-hover);
  transform: translateY(-1px);
}

.sidebar {
  width: 40px;
  height: calc(100vh - 50px);

  padding-top: 10px;
  list-style: none;
  background-color: var(--color-background-header);
  border-right: 1px solid var(--color-border-default);
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.05);

  display: flex;
  gap: 5px;
  flex-direction: column;
  align-items: center;

  position: relative;

  z-index: 99999;

  transition: all 0.3s ease;
}

.sidebar::before {
  content: "";
  width: 35px;
  height: 35px;

  background-color: var(--color-accent);
  border-radius: 8px;

  position: absolute;
  top: v-bind(sidebar_before_top);

  transition: all 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
  box-shadow: 0 2px 8px rgba(66, 184, 131, 0.3);
}

.sidebar:deep(svg) {
  fill: var(--color-text-header);
  transition: all 0.2s ease;
}

.sidebar > li {
  width: 35px;
  height: 35px;

  border-radius: 8px;

  position: relative;

  display: flex;
  align-items: center;
  justify-content: center;

  cursor: pointer;

  transition: all 0.2s ease;
  z-index: 2;
}

.sidebar > li:hover {
  background: linear-gradient(135deg, var(--color-accent), rgba(66, 184, 131, 0.3));
  transform: scale(1.15);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  border-radius: 10px;
}

.sidebar-text {
  width: auto;
  min-width: 50px;
  height: 20px;

  margin: 0;
  padding: 5px 10px;
  box-sizing: border-box;
  border-radius: 4px;
  font-size: 13px;
  color: var(--color-text-header);
  background-color: var(--color-background-header);

  position: absolute;
  left: 40px;

  display: flex;
  justify-content: flex-start;
  align-items: center;
  white-space: nowrap;

  opacity: 0;
  transform: translateX(-55px);

  transition: opacity 0.25s ease, transform 0.3s ease;
}

.sidebar > li:hover .sidebar-text {
  opacity: 1;
  transform: translateX(0);
}

.view-box {
  width: calc(100% - 40px);
  height: calc(100% - 50px);

  overflow: hidden;

  position: absolute;
  right: 0;
  bottom: 0;
}

.fade-enter-active,
.fade-leave-active {
  transition: all 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
  position: absolute;
  width: 100%;
  top: 0;
}

.fade-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.fade-enter-to {
  opacity: 1;
  transform: translateX(0);
}

.fade-leave-from {
  opacity: 1;
  transform: translateX(0);
}

.fade-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}
</style>