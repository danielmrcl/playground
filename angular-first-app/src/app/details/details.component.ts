import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HousingService } from '../housing.service';
import { HousingLocation } from '../housing-location';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';

@Component({
  standalone: true,
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.scss'],
  imports: [
    ReactiveFormsModule
  ]
})
export class DetailsComponent {
  readonly housingService: HousingService;

  housingLocation: HousingLocation | undefined;

  applyForm = new FormGroup({
    firstName: new FormControl(),
    lastName: new FormControl(),
    email: new FormControl()
  })

  constructor(route: ActivatedRoute, housingService: HousingService) {
    this.housingService = housingService;
    this.housingLocation =
      housingService.getHousingLocationById(Number(route.snapshot.params['id']));
  }

  submitApplication() {
    this.housingService.submitApplication(
      this.applyForm.value.firstName ?? '',
      this.applyForm.value.lastName ?? '',
      this.applyForm.value.email ?? ''
    );
  }
}
