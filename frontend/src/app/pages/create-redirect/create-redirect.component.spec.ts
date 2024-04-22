import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateRedirectComponent } from './create-redirect.component';

describe('CreateRedirectComponent', () => {
  let component: CreateRedirectComponent;
  let fixture: ComponentFixture<CreateRedirectComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CreateRedirectComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(CreateRedirectComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
