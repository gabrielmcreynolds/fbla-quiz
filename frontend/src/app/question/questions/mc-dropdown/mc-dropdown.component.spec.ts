import { ComponentFixture, TestBed } from '@angular/core/testing';

import { McDropdownComponent } from './mc-dropdown.component';

describe('McDropdownComponent', () => {
  let component: McDropdownComponent;
  let fixture: ComponentFixture<McDropdownComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ McDropdownComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(McDropdownComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
