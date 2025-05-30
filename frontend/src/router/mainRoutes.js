import MainLayout from "@/layouts/MainLayout.vue";
import ChartView from "@/views/ChartView.vue";
import IndexView from "@/views/IndexView.vue";


const routes = [
    {
        path: "/",
        component: MainLayout,
        children: [
            {
                path: "/",
                name: "Index",
                component: IndexView,
            },
            {
                path: "/chart",
                name: "Chart",
                component: ChartView
            }
        ],
    },
];

export default routes;
