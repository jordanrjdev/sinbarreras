import express from "express";
import IndexRoutes from "./routes/index.routes";

const app: express.Application = express();
const PORT = 4000;

app.use(express.json());
app.use(IndexRoutes);

app.listen(PORT, () => {
  console.log(`Express with Typescript! http://localhost:${PORT}`);
});
