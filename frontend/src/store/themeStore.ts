import { defineStore } from 'pinia'

type ThemeFlavor = 'catppuccin-latte' | 'catppuccin-frappe' | 'catppuccin-macchiato' | 'catppuccin-mocha'

const VALID_THEMES: ThemeFlavor[] = ['catppuccin-latte', 'catppuccin-frappe', 'catppuccin-macchiato', 'catppuccin-mocha']

export const useThemeStore = defineStore('theme', {
  state: () => ({ currentTheme: 'catppuccin-mocha' as ThemeFlavor }),
  getters: {
    isDark(): boolean { return this.currentTheme !== 'catppuccin-latte' },
    label(): string {
      return {
        'catppuccin-latte': 'Latte',
        'catppuccin-frappe': 'Frappé',
        'catppuccin-macchiato': 'Macchiato',
        'catppuccin-mocha': 'Mocha'
      }[this.currentTheme]
    }
  },
  actions: {
    applyTheme(flavor: ThemeFlavor) {
      this.currentTheme = flavor
      document.documentElement.setAttribute('data-theme', flavor)
      document.documentElement.setAttribute('data-bulma-theme', flavor === 'catppuccin-latte' ? 'light' : 'dark')
    },
    init() {
      if (VALID_THEMES.includes(this.currentTheme)) {
        this.applyTheme(this.currentTheme)
        return
      }
      const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
      this.applyTheme(prefersDark ? 'catppuccin-mocha' : 'catppuccin-latte')
    }
  },
  persist: {
    storage: localStorage,
    pick: ['currentTheme'],
  }
})
