const defaultTheme = require("tailwindcss/defaultTheme")
const plugin = require("tailwindcss/plugin")

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.templ"],
  theme: {
    extend: {
      colors: {
        rwhite: "#ffffff",
        rblack: "#141414",
        rdarker: "#1b5b57",
        rdark: "#257f7a",
        rprimary: "#35b5ae",
        rlight: "#72cbc6",
        rlighter: "#aee1df",
        rlightest: "#ebf8f7",
        rgray: "#cccccc",
        rgraylightest: "#f1f5f9",
        rgraylighter: "#e1e1e1",
        rgraylight: "#dddddd",
        rgraydark: "#bababa",
        rgraydarker: "#7e7e7e",
        rgraydarkest: "#434343",
        rred: "#ef4444",
        rredlight: "#ff9494",
        rgreen: "#06ff00",
        rgreenlight: "#9ad4a9",
        rgreendark: "#34a853",
        ryellow: "#fabb05",
      },
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
    plugin(function ({ addUtilities }) {
      addUtilities({
        ".text-last-center": {
          "textAlignLast": "center",
        },
      })
    })
  ],
}

