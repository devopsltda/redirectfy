import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DevopsBrandComponent } from './devops-brand.component';

describe('DevopsBrandComponent', () => {
  let component: DevopsBrandComponent;
  let fixture: ComponentFixture<DevopsBrandComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DevopsBrandComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(DevopsBrandComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
