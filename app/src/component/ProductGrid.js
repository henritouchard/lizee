import React, { useState, useEffect } from "react";

import Paper from "@material-ui/core/Paper";
import Tabs from "@material-ui/core/Tabs";

import QuantityPicker from "./QuantityPicker";
import { ProductCard, FlexContainer, TabPanel, StyledTab } from "./styles/Containers";
import Cart from "./Cart";

function a11yProps(index) {
  return {
    id: `full-width-tab-${index}`,
    "aria-controls": `full-width-tabpanel-${index}`,
  };
}

function ProductGrid({ products, addToCart, cartProducts, dates }) {
  const [newProduct, setNewProduct] = useState(null);
  const [selectedTab, setSelectedTab] = useState(0);

  const handleNewProducts = (product, quantity) => {
    let existing = false;
    let cp = cartProducts;
    cp.forEach((element) => {
      if (element.product.id === product.id) {
        element.quantity += quantity;
        existing = true;
      }
    });
    if (existing) addToCart(cp);
    else {
      addToCart([
        ...cartProducts,
        {
          product,
          quantity,
          from: dates.from,
          to: dates.to,
        },
      ]);
    }
  };

  const handleChange = (event, tab) => {
    setSelectedTab(tab);
  };

  // Update cart when needed
  useEffect(() => {
    if (newProduct != null) {
      addToCart([...cartProducts, newProduct]);
      setNewProduct(null);
    }
  }, [newProduct, addToCart, cartProducts]);

  const inCart = cartProducts;
  return (
    <>
      <Paper square>
        <Tabs
          value={selectedTab}
          indicatorColor="primary"
          textColor="primary"
          onChange={handleChange}
          aria-label="products-Cart switcher"
          variant="fullWidth"
          TabIndicatorProps={{ style: { background: "rgba(0,0,0,0)" } }}
        >
          <StyledTab active={selectedTab === 0 ? 1 : 0} label="Shop" {...a11yProps(0)} />
          <StyledTab
            active={selectedTab === 1 ? 1 : 0}
            label={`Cart(${inCart.length})`}
            disabled={!cartProducts || cartProducts.length === 0}
            {...a11yProps(1)}
          />
        </Tabs>
      </Paper>

      <TabPanel value={selectedTab} index={0}>
        <FlexContainer>
          {products.error ? (
            <h3 style={{ display: "block", color: "red", margin: "0 auto" }}>
              {products.error}
            </h3>
          ) : (
            products.map((product) => {
              const { id, name, picture, availability } = product;
              return (
                <li key={id}>
                  <ProductCard>
                    <h3>{name}</h3>
                    <img src={picture} style={{ width: "200px" }} alt={name}></img>
                    <QuantityPicker product={product} addToCart={handleNewProducts} />
                    <p style={{ fontSize: "12px", color: "grey" }}>
                      In stock : {availability}
                    </p>
                  </ProductCard>
                </li>
              );
            })
          )}
        </FlexContainer>
      </TabPanel>

      <TabPanel value={selectedTab} index={1}>
        <Cart dates={dates} addToCart={addToCart} cartProducts={cartProducts}></Cart>
      </TabPanel>
    </>
  );
}

export default ProductGrid;
