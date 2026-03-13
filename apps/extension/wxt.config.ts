import { defineConfig } from 'wxt';

// See https://wxt.dev/api/config.html
export default defineConfig({
  modules: ['@wxt-dev/module-react'],
  manifest: {
    name: 'SceneShare',
    description: 'Share a moment from any video',
    permissions: ['activeTab', 'clipboardWrite', 'storage'],
    host_permissions: [
      'https://www.youtube.com/*',
      'http://localhost:4006/*',
    ]
  }
});
