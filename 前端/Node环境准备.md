# Node.js & npm 安装配置指南

## 1. 安装Node.js和npm
```bash
# 使用Homebrew安装（推荐）
brew install node

# 直接下载安装（可选1）
curl -o node-v22.15.0.pkg https://nodejs.org/dist/v22.15.0/node-v22.15.0.pkg && sudo installer -pkg node-v22.15.0.pkg -target /

# 通过 nvm 安装（可选2）
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.2/install.sh | bash
\. "$HOME/.nvm/nvm.sh"
nvm install 22
nvm current

# 验证安装
node -v
npm -v
```

## 2. 配置国内镜像源
```bash
# 设置淘宝镜像（推荐）
npm config set registry https://registry.npmmirror.com

# 或配置腾讯云镜像
npm config set registry https://mirrors.cloud.tencent.com/npm/

# 验证配置
npm get registry
```

## 3. 常用工具安装
```bash
# 使用国内镜像安装yarn
npm install -g yarn --registry=https://registry.npmmirror.com

# 安装pnpm
npm install -g pnpm --registry=https://registry.npmmirror.com
```

## 4. 临时使用镜像
```bash
# 单次安装使用镜像
npm install --registry=https://registry.npmmirror.com
```

## 5. 恢复默认配置

## 6. 前端开发工具链

### 6.1 VS Code 配置
```bash
# 安装推荐扩展
code --install-extension dbaeumer.vscode-eslint
code --install-extension esbenp.prettier-vscode
code --install-extension Vue.volar
code --install-extension stylelint.vscode-stylelint

# 创建工作区配置（.vscode/settings.json）
{
  "editor.formatOnSave": true,
  "editor.defaultFormatter": "esbenp.prettier-vscode",
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": true
  },
  "typescript.tsdk": "node_modules/typescript/lib"
}
```

### 6.2 脚手架使用
```bash
# 使用Vite创建项目（Vue示例）
npm create vite@latest my-vue-app -- --template vue
cd my-vue-app
npm install
npm run dev

# 使用Create React App
npx create-react-app my-react-app
cd my-react-app
npm start

# 项目构建
npm run build
```

## 7. 现代前端框架对比
| 特性        | Vue 3                  | React 18               |
|-----------|------------------------|-----------------------|
| 学习曲线      | 渐进式，模板语法更易上手       | JSX需要适应，概念相对更多     |
| 状态管理      | Pinia                 | Redux/Context         |
| 样式方案      | 单文件组件样式隔离           | CSS-in-JS方案         |
| 类型支持      | 原生TS支持               | 需要额外配置类型声明          |
| 适用场景      | 快速原型/中等规模应用         | 复杂企业级应用/跨平台开发       |

## 8. 浏览器调试技巧
1. Chrome DevTools 快捷键:
   - Ctrl+Shift+C 快速选择元素
   - Ctrl+Shift+P 打开命令面板
2. 网络请求调试:
   - 查看请求头/响应头
   - 模拟慢速网络（Network > Online）
3. 性能分析:
   - Performance面板录制运行时性能
   - Lighthouse生成优化报告

推荐插件：
- React Developer Tools
- Vue.js devtools
- Redux DevTools
- JSON Formatter
```bash
# 清除所有镜像配置
npm config clean

# 或手动删除配置
rm ~/.npmrc
```

**注意事项**：
1. 镜像配置会写入`~/.npmrc`文件
2. 推荐同时设置`ELECTRON_MIRROR`环境变量加速Electron下载
3. 使用yarn时需单独配置镜像：
```bash
yarn config set registry https://registry.npmmirror.com
```

        