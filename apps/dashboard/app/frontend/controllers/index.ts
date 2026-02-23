import { application } from "./application"
import { definitionsFromGlobImports } from "stimulus-vite-helpers"

const controllers = import.meta.glob("./**/*_controller.ts", { eager: true })
application.load(definitionsFromGlobImports(controllers))
