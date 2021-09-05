import { Router } from "express";
import {
  getAvatars,
  helloWorld,
  getAvatarbyId,
} from "../controllers/index.controller";

const router: Router = Router();

router.route("/").get(helloWorld);

router.route("/avatar/:id").get(getAvatarbyId);

router.route("/avatars").get(getAvatars);

export default router;
