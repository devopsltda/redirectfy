import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FormCreateRedirectComponent } from './form-create-redirect.component';

describe('FormCreateRedirectComponent', () => {
  let component: FormCreateRedirectComponent;
  let fixture: ComponentFixture<FormCreateRedirectComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FormCreateRedirectComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(FormCreateRedirectComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
