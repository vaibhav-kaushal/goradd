/**
 * testsuite contains the unit tests that test the goradd.js javascript file, in conjunction with the jsunit_form go
 * form. Tests start with "test".
 */
goradd.testsuite = {
    /**
     * Reset the test space to its original state
     * @param t
     */
    pretest: function(t) {
        goradd.el("testspace").innerHTML =  '<p id="testP">I am here</p> \
        <div id="testD" data-animal-type="bird" spellcheck>a div</div> \
        <input id="testC" type="checkbox">';
        //        <ul id="listener"><li id="outer"><span id="inner">a span</span></li></ul>
    },
    testEl: function(t) {
       var el = goradd.el("testP");
       t.assert(el.innerText === "I am here");
    },
    testQS: function(t) {
        var el = goradd.qs("p[id='testP']");
        t.assert(el.innerText === "I am here");

        el = g$("testspace").qs("p[id='testP']");
        t.assert(el.innerText === "I am here");
    },
    testQA: function(t) {
        var el = goradd.qa("p[id='testP']");
        t.assert(el[0].innerText === "I am here");

        el = g$("testspace").qa("p[id='testP']");
        t.assert(el[0].innerText === "I am here");
    },
    testIsEmptyObj: function(t) {
        t.assert(!goradd.isEmptyObject({a:"b"}));
        t.assert(goradd.isEmptyObject({}));
    },
    testForm: function(t) {
        t.isSame("JsUnitTestForm", goradd.form().id);
    },
    testMatches: function(t) {
        t.assert(g$("testP").matches("p[id='testP']"));
        t.assert(!g$("testP").matches("div"));
    },
    testParents: function(t) {
        var p = g$("testP").parents();
        t.isSame("JsUnitTestForm", p[1].id);
    },
    testClosest: function(t) {
        var p = g$("testP");
        var c = p.closest("form");
        t.isSame("JsUnitTestForm", c.element.id);
    },
    testAttrProp: function(t) {
        var a1 = g$("testD").attr("spellcheck");
        t.isSame(true, a1);
        g$("testD").prop({"spellcheck":false, "class":"a"});
        t.isSame(false, g$("testD").prop("spellcheck"));
        t.isSame("a", g$("testD").prop("class"));
        g$("testD").prop("class", "b");
        t.isSame("b", g$("testD").prop("class"));
    },
    testClass: function(t) {
        g$("testD").class("b c");
        g$("testD").class("-c");
        t.isSame("b", g$("testD").class());

        g$("testD").class("+a");
        g$("testD").class("-b");
        t.isSame("a", g$("testD").class());
    },
    testEvent: function(t) {
        g$("testD").on("et", function() {
            this.element.innerText = "tested";
        });
        g$("testD").trigger("et");
        t.isSame("tested", goradd.el("testD").innerText)

        g$("listener").on("click", "li", function(event) {
            t.isSame("listener", this.element.id);
            t.isSame("listener", event.currentTarget.id);
            t.isSame("inner", event.target.id);
            t.isSame("outer", event.goradd.match.id);
        }, {bubbles: true});
        g$("inner").click();
    },
    testHtmlInserts: function(t) {
        var p = g$("testP");
        p.htmlAfter("<p id='after'>Inserted After</p>");
        p.htmlBefore("<p id='before'>Inserted Before</p>");
        t.isSame("Inserted After", goradd.el("after").innerText);
        t.isSame("Inserted Before", goradd.el("before").innerText);
        goradd.el("testP").innerText = "There";
        p.insertChildHtml("Here");
        p.appendChildHtml("Everywhere");
        t.isSame("HereThereEverywhere", goradd.el("testP").innerText);
    },
    testMoves: function(t) {
        var p = g$("testP");
        p.moveAfter("testD");
        var d = p.prev();
        t.isSame("testD", d.id, "moveAfter failed");

        p.moveBefore(d);
        var n = p.next();
        t.isSame("testD", n.id, "moveAfter failed");

        p.htmlAfter("<div id='after'><div id='in'>I am in.</div></div>");
        p.insertInto("after");
        n = p.next();
        t.isSame("in", n.id, "insertInto failed");

        p.appendTo("after");
        d = p.prev();
        t.isSame("in", d.id, "appendTo failed");
    },
    testRemove: function(t) {
        g$("testP").remove();
        t.isSame(goradd.el("testP"), null);
    },
    testEach: function(t) {
        var s = "";
        goradd.each(g$("testspace").qa("p,div"), function(i, v) {
            s += v.innerText;
        });
        t.isSame("I am herea div", s);

        s = "";
        goradd.each(["a", "b", "c"], function(i,v) {
            s += v;
        });
        t.isSame("abc", s);

        s = [];
        goradd.each({a:"a", b:"b", c:"c"}, function(k,v) {
            s.push(k+v);
        });
        t.assert(s.indexOf("aa") !== -1);
    },
    testData: function(t) {
        var d = g$("testD");
        t.isSame("bird", d.data("animalType"));
        d.data("animalType", "dog");
        t.isSame("dog", d.data("animalType"));
    },
    testClick: function(t) {
        // Here we are making sure that our generated click function will also trigger subsequent events
        // that go along with a click
        var c = g$("testC");
        var p = goradd.el("testP");
        p.innerText = "";
        c.on("click", function(e) {
            p.innerText = p.innerText + "clicked";
        });
        c.on("change", function(e) {
            p.innerText = p.innerText + "changed";
        });
        c.click();
        t.isSame("clickedchanged", p.innerText);
    }
};
