import React, { useCallback, useMemo, useState } from "react";
import { AngleDown } from "../icons/AngleDown";

export const Dropdown = ({
  className,
  placeholder,
  value,
  options,
  onChange,
  ...props
}) => {
  const [show, setShow] = useState(false);
  const [selected, setSelected] = useState(
    options ? options.find((o) => o.value === value) : null
  );
  const setValue = useCallback(
    (val, option) => {
      !!onChange && onChange(val, option);
      setSelected(option);
      setShow(false);
    },
    [value, onChange, setShow]
  );

  return (
    <div
      {...props}
      className={`dropdown ${className || ""} ${show ? "is-active" : ""}`}
    >
      <div className="dropdown-trigger w-full">
        <button
          type="button"
          className="button is-flex is-justify-content-space-between w-full"
          aria-haspopup="true"
          onClick={() => setShow((v) => !v)}
        >
          <span>
            {(selected ? selected.label || selected.value : null) ||
              placeholder ||
              "Select Item"}
          </span>
          <AngleDown className="icon is-small is-size-7" />
        </button>
      </div>
      <div className="dropdown-menu w-full" role="menu">
        <div
          className="dropdown-content"
          style={{ maxHeight: "200px", overflowY: "scroll" }}
        >
          {!!options &&
            options.map((option) => (
              <a
                key={option.value}
                href="#"
                className="dropdown-item"
                onClick={() => setValue(option.value, option)}
              >
                {option.label || option.value}
              </a>
            ))}
        </div>
      </div>
    </div>
  );
};
