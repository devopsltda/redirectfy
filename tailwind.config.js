const defaultTheme = require("tailwindcss/defaultTheme")
const plugin = require("tailwindcss/plugin")
const Color = require("color")

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.templ"],
  theme: {
    colors: {
      transparent: 'transparent',
      rwhite: "#ffffff",
      rblack: "#141414",
      rdarker: "#1b5b57",
      rdark: "#257f7a",
      rprimary: "#35b5ae",
      rlight: "#72cbc6",
      rlighter: "#aee1df",
      rlightest: "#ebf8f7",
      rgray: "#cccccc",
      rgraylightest: "#f2f2f2",
      rgraylighter: "#e1e1e1",
      rgraylight: "#dddddd",
      rgraydark: "#bababa",
      rgraydarker: "#7e7e7e",
      rgraydarkest: "#434343",
      rred: "#ff0000",
      rredlight: "#ff9494",
      rgreen: "#06ff00",
      rgreenlight: "#9ad4a9",
      rgreendark: "#34a853",
      ryellow: "#fabb05",
    },
    extend: {
      maxHeight: {
        '0': '0',
        xl: '36rem',
      },
      fontFamily: {
        'sans': ['"DM Sans"', ...defaultTheme.fontFamily.sans],
      },
    },
  },
  variants: {
    backgroundColor: [
      'hover',
      'focus',
      'active',
      'odd',
    ],
    display: ['responsive'],
    textColor: [
      'focus-within',
      'hover',
      'active',
    ],
    placeholderColor: ['focus'],
    borderColor: ['focus', 'hover'],
    boxShadow: ['focus'],
  },
  plugins: [
    plugin(({ addUtilities, e, theme, variants }) => {
      const newUtilities = {}
      Object.entries(theme('colors')).map(([name, value]) => {
        if (name === 'transparent' || name === 'current') return
        const color = value[300] ? value[300] : value
        const hsla = Color(color).alpha(0.45).hsl().string()

        newUtilities[`.shadow-outline-${name}`] = {
          'box-shadow': `0 0 0 3px ${hsla}`,
        }
      })

      addUtilities(newUtilities, variants('boxShadow'))
    }),
    plugin(function ({ addUtilities }) {
      addUtilities({
        ".text-last-center": {
          "textAlignLast": "center",

        },
      })
    })
  ],
}

