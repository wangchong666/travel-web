import { NgModule } from '@angular/core';
import { AngularEchartsModule } from 'ngx-echarts';

import { ThemeModule } from '../../@theme/theme.module';
import { PostsComponent } from './posts.component';
// import { NgxDatatableModule } from '@swimlane/ngx-datatable';
import { Ng2SmartTableModule, } from 'ng2-smart-table';
import { FroalaEditorModule, FroalaViewModule } from 'angular-froala-wysiwyg';


@NgModule({
  imports: [
    ThemeModule,
    AngularEchartsModule,
    // NgxDatatableModule,
    Ng2SmartTableModule,
    FroalaEditorModule.forRoot(), 
    FroalaViewModule.forRoot()
  ],
  declarations: [
    PostsComponent
  ],
})
export class PostsModule { }
