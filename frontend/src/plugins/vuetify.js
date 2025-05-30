import "vuetify/styles";
import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import { aliases, mdi } from "vuetify/iconsets/mdi";
import { myTheme } from "./theme";

export const vuetify = createVuetify({
    components,
    directives,
    theme: {
        defaultTheme: "myTheme",
        themes: {
            myTheme,
        },
    },
    icons: {
        defaultSet: "mdi",
        aliases,
        sets: {
            mdi,
        },
    },
});
