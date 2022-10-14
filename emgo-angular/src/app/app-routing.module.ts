import {NgModule} from '@angular/core';
import { RouterModule, Routes} from "@angular/router";

const routes: Routes = [
  {path: 'subjects', loadChildren: () => import('./subjects/subjects.module').then(m => m.SubjectsModule)}
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})

export class AppRoutingModule {}
