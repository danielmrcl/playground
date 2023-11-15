import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'angular-first-app';
  showMessages = true;
  message = 'Hello world!!!!';

  toggleMessages() {
    this.showMessages = !this.showMessages;
  }
}
