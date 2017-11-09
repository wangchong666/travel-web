import { Component } from '@angular/core';
import { ServerDataSource } from 'ng2-smart-table';
import { Http } from '@angular/http';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'ngx-posts',
  styleUrls: ['./posts.component.scss'],
  templateUrl: './posts.component.html',
})
export class PostsComponent {
  settings = {
    add: {
      addButtonContent: '<i class="nb-plus"></i>',
      createButtonContent: '<i class="nb-checkmark"></i>',
      cancelButtonContent: '<i class="nb-close"></i>',
    },
    edit: {
      editButtonContent: '<i class="nb-edit"></i>',
      saveButtonContent: '<i class="nb-checkmark"></i>',
      cancelButtonContent: '<i class="nb-close"></i>',
    },
    delete: {
      deleteButtonContent: '<i class="nb-trash"></i>',
      confirmDelete: true,
    },
    mode: 'external',
    // pager:{
    //   display : true,
    //   perPage : 5
    // },
    columns: {

      Name: {
        title: 'Name',
        type: 'string',
      },
      Category: {
        title: 'Category',
        type: 'string',
      },
      Address: {
        title: 'Address',
        type: 'string',
      },
      Site: {
        title: 'Site',
        type: 'string',
      },
      Publish: {
        title: 'Publish',
        type: 'bool',
      },
    },
  };

  source: ServerDataSource //= new ServerDataSource();

  // constructor(private service: SmartTableService) {
  //   const data = this.service.getData();
  //   this.source.load(data);
  // }

  constructor(http: Http, private modalService: NgbModal) {
    this.source = new ServerDataSource(http, { endPoint: 'http://localhost:8083/api/posts', dataKey: 'Data', totalKey: 'Total' });
  }

  // constructor(private modalService: NgbModal) { }

  currentData:any={}
  showEdit(content) {
    const activeModal = this.modalService.open(content, { size: 'lg', container: 'nb-layout'});
    activeModal.result.then((result) => {
      // this.closeResult = `Closed with: ${result}`;
    }, (reason) => {
      // this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
    });
  }


    // activeModal.componentInstance.modalHeader = 'Large Modal';
  

  // private getDismissReason(reason: any): string {
  //   if (reason === ModalDismissReasons.ESC) {
  //     return 'by pressing ESC';
  //   } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
  //     return 'by clicking on a backdrop';
  //   } else {
  //     return  `with: ${reason}`;
  //   }
  // }

  onDeleteConfirm(event): void {
    if (window.confirm('Are you sure you want to delete?')) {
      event.confirm.resolve();
    } else {
      event.confirm.reject();
    }
  }

  onEdit(event,content) {
    Object.assign(this.currentData,event.data)
    if(!this.currentData.Content.cn){
      this.currentData.Content.cn = {}
      Object.assign(this.currentData.Content.cn,this.currentData.Content.en)
    }
    this.objectKeys = Object.keys(this.currentData.Content)

    this.showEdit(content)
  }

  objectKeys =[]
}
