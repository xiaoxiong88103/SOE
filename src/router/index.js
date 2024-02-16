import { createBrowserRouter } from "react-router-dom";
import RouterDate from "./router";
const router = createBrowserRouter([...RouterDate.Over, RouterDate.Index]);
export default router;
